package main

import (
	"errors"
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/sirupsen/logrus"
	"os"
)

type Hook struct {
	id string
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	return nil
}

func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func ThrowError() error {
	err := errors.New("a crazy failure")
	log.WithField("id", "123").Trace("ThrowError").Stop(&err)
	return err
}

type CustomHandler struct {
	id      string
	handler log.Handler
}

func (h *CustomHandler) HandleLog(e *log.Entry) error {
	e.WithField("id", h.id)
	return h.handler.HandleLog(e)
}

func main() {
	// Logrus
	fmt.Println("Logrus:")

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "Just now"}

	x := logrus.WithFields(fields)
	x.Warn("warning!")
	x.Error("error!")
	fmt.Println()

	// Apex
	fmt.Println("Apex:")
	log.SetHandler(&CustomHandler{"123", text.New(os.Stdout)})
	err := ThrowError()

	log.WithError(err).Error("an error occurred")
}
