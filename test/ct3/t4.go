package ct3

import (
	"encoding/json"
	"fmt"
)

// https//wookiist.dev/93  json  형식에 자주 사용
const (
	jsonBlob     = `{"name": "Aaron"}`
	fulljsonBlob = `{"name":"Aaron", "age":0}`
)

// 구조체
type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

//인코딩 디코딩
func T4() {

	if err := baseEncoding(); err != nil {
		panic(err)
	}
}

//마샬링 언마샬링 https://jeonghwan-kim.github.io/dev/2019/01/18/go-encoding-json.html
func baseEncoding() error {
	e := Example{}
	// note that no age = 0 age
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with age = 0:", string(value))

	return nil
}
