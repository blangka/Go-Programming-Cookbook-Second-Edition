package ct4

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func originalError() error {
	return errors.New("error occurred")
}

func passThroughError() error {
	err := originalError()
	return errors.Wrap(err, "in passthrougherror")
}

func T3() {
	fmt.Println("basic logging and modification of logger:")
	buf := bytes.Buffer{}

	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)

	logger.Println("test")

	logger.SetPrefix("new logger: ")

	logger.Printf("you can also add args(%v) and use Fatalln to log and crash", true)

	fmt.Println("+" + buf.String())

	fmt.Println("logging 'handled' errors:")

	err := passThroughError()
	if err != nil {
		// we log because an unexpected error occurred!
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
