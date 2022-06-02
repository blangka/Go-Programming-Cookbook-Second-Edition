package main

import (
	"encoding/json"
	"fmt"
	"github.com/blangka/Go-Programming-Cookbook-Second-Edition/ct8/t2"
	"net/http"
)

func main() {
	storage := t2.MemStorage{}
	c := new(&storage)
	http.HandleFunc("/set", c.setValue)
	http.HandleFunc("/get", c.getValue(false))
	http.HandleFunc("/get/default", c.getValue(true))

	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}

// controller

type controller struct {
	storage t2.Storage
}

func new(storage t2.Storage) *controller {
	return &controller{
		storage: storage,
	}
}

type payload struct {
	Value string `json:"value"`
}

//set get
func (c *controller) setValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	value := r.FormValue("value")
	c.storage.Put(value)
	w.WriteHeader(http.StatusOK)
	p := payload{Value: value}
	if payload, err := json.Marshal(p); err == nil {
		w.Write(payload)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (c *controller) getValue(UseDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		value := "default"
		if !UseDefault {
			value = c.storage.Get()
		}
		p := payload{Value: value}
		w.WriteHeader(http.StatusOK)
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}
	}
}
