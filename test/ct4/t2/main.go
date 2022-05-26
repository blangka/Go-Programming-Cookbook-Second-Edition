package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

type ErrorTyped struct {
	error
}

func main() {
	// Wrap
	e := errors.New("standard error")
	fmt.Println("Regular Error - ", WrappedError(e))
	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))
	fmt.Println("Nil -", WrappedError(nil))
	fmt.Println()

	// Unwrap
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Println("wrapped error: ", err)

	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occurred")
	}
	fmt.Println()

	// stack trace
	err = error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}
