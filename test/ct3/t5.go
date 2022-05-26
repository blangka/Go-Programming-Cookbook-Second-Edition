package ct3

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type pos struct {
	X      int
	Y      int
	Object string
}

// 인코딩 디코딩
func T5() {

	if err := gobEx(); err != nil {
		fmt.Println(err)
	}

	if err := base64Ex(); err != nil {
		fmt.Println(err)
	}
}

func gobEx() error {
	buffer := bytes.Buffer{}

	p := pos{
		X:      10,
		Y:      15,
		Object: "lch",
	}

	// p가 인터페이스 인경우 gob.NewEncoder를 호출해야 한다.
	e := gob.NewEncoder(&buffer)
	if err := e.Encode(&p); err != nil {
		return err
	}
	fmt.Println("Gob Encode value : ", buffer.Bytes())

	p2 := pos{}
	d := gob.NewDecoder(&buffer)
	if err := d.Decode(&p2); err != nil {
		return err
	}

	fmt.Println("Gob Decode value : ", p2)

	return nil
}

func base64Ex() error {

	value := base64.URLEncoding.EncodeToString([]byte("encoding some data!"))
	fmt.Println("With EncodeToString and URLEncoding: ", value)

	// decode the first value
	decoded, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	fmt.Println("With DecodeToString and URLEncoding: ", string(decoded))

	return nil
}
