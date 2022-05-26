package ct3

import "fmt"

//클로져
// https://hwan-shell.tistory.com/339
// http://golang.site/go/article/11-Go-%ED%81%B4%EB%A1%9C%EC%A0%80
func T7() {

	next := nextValue()

	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	anotherNext := nextValue()
	fmt.Println(anotherNext()) // 1 다시 시작
	fmt.Println(anotherNext()) // 2

}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
