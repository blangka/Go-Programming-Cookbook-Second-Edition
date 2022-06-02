package main

import (
	"fmt"
	"github.com/blangka/Go-Programming-Cookbook-Second-Edition/test/ct8/t2"
	"net/http"
)

// Controller passes state to our handlers
type controller struct {
	storage Storage
}

// New is a Controller 'constructor'
func New(storage Storage) *controller {
	return &controller{
		storage: storage,
	}
}

// Payload is our common response
type payload struct {
	Value string `json:"value"`
}

func main() {
	storage := controllers.MemStorage{}
	c := controllers.New(&storage)
	http.HandleFunc("/get", c.GetValue(false))
	http.HandleFunc("/get/default", c.GetValue(true))
	http.HandleFunc("/set", c.SetValue)

	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
