package ct3

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Name string `serialize:"name"`
	City string `serialize:"city"`
	Year int    `serialize:"year"`
}

// 리플렉트
func T6() {
	emptySt()
	notEmptySt()
}

func emptySt() {
	p := Person{}

	res, err := serializeStructStrings(&p)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Empty struct: %#v\n", p)
	fmt.Printf("Serialize Result:", res)

	newP := Person{}
	if err := deSerializeStructStrings(res, &newP); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Deserialize results: %#v\n", newP)
}

func notEmptySt() {
	p := Person{
		Name: "lim",
		City: "suji",
		Year: 2017,
	}

	res, err := serializeStructStrings(&p)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("full struct: %#v\n", p)
	fmt.Printf("Serialize Result:", res)

	newP := Person{}
	if err := deSerializeStructStrings(res, &newP); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Deserialize results: %#v\n", newP)
}

func serializeStructStrings(s interface{}) (string, error) {
	result := ""

	r := reflect.TypeOf(s)
	value := reflect.ValueOf(s)

	//포인터 인경우 처리
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		key := field.Name
		if serialize, ok := field.Tag.Lookup("serialize"); ok {

			key = serialize
		}

		switch value.Field(i).Kind() {
		case reflect.String:
			result += key + ":" + value.Field(i).String() + ";"
		default:
			continue
		}
	}

	return result, nil
}

func deSerializeStructStrings(s string, res interface{}) error {
	r := reflect.TypeOf(s)

	if r.Kind() != reflect.Ptr {
		return errors.New("res must pointer")
	}

	r = r.Elem()
	value := reflect.ValueOf(res).Elem()

	vals := strings.Split(s, ";")
	valMap := make(map[string]string)
	for _, v := range vals {
		keyval := strings.Split(v, ":")
		if len(keyval) != 2 {
			continue
		}
		valMap[keyval[0]] = keyval[1]
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		if serialize, ok := field.Tag.Lookup("serialize"); ok {

			if val, ok := valMap[serialize]; ok {
				value.Field(i).SetString(val)
			}
		} else if val, ok := valMap[field.Name]; ok {
			value.Field(i).SetString(val)
		}
	}

	return nil
}
