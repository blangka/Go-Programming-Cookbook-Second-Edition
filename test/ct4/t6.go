package ct4

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log     *logrus.Logger
	initLog sync.Once
)

func useLog() error {
	err := errors.New("already initialized")
	initLog.Do(func() {
		err = nil
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
	})
	if err != nil {
		return err
	}

	log.WithField("key", "value").Debug("hello")
	log.Debug("test")

	return nil
}

func T6() {
	if err := useLog(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := useLog(); err != nil {
		panic(err)
	}
}
