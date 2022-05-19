package ct3

import "fmt"

// 형변환
func T1() {

	// int
	var a = 24

	// float 64
	var b = 2.0
	c := float64(a) * b
	fmt.Println(c)

	var i interface{}
	i = "test"

	if val, ok := i.(string); ok {
		fmt.Println("string 타입의 값은", val)
	}

	if _, ok := i.(int); !ok {
		fmt.Println("int가 아닌 값이다.")
	}

	//http://golang.site/go/article/7-Go-%EC%A1%B0%EA%B1%B4%EB%AC%B8 참고
}
