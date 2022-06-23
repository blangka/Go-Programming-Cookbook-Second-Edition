package main

import (
	"fmt"
	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter8/validation"
	"github.com/blangka/Go-Programming-Cookbook-Second-Edition/ct8/t3"
	"net/http"
)

func main() {
	c := validation.New()
	http.HandleFunc("/", c.Process)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
