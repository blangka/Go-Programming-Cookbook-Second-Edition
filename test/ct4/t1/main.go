package main

import (
	"errors"
	"fmt"
)

var ErrorValue = errors.New("this is a typed error")

type TypedError struct {
	error
}

type CustomError struct {
	Result string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("there was an error; %s was the result", c.Result)
}

func main() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorValue
	fmt.Println("value error: ", err)

	err = TypedError{errors.New("typed error")}
	fmt.Println("typed error: ", err)

	err = CustomError{Result: "this"}
	fmt.Println("custom error: ", err)
}
