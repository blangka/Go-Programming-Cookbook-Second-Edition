package ct3

import (
	"fmt"
	"math"
	"math/big"
)

//http://golang.site/go/article/14-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Map
var memorize map[int]*big.Int

//https://blog.realsangil.net/posts/2018-03-13-go-init-function/
func init() {
	//메모리 초기화
	memorize = make(map[int]*big.Int)
}

// math 를 이용한 숫자 데이터 타입
func T2() {
	i := 10
	fmt.Println(fib(i))

	fmt.Println(math.Ceil(9.5))
	fmt.Println(math.Floor(9.5))
}

func fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}

	// base case
	if n < 2 {
		memorize[n] = big.NewInt(1)
	}
	memorize[n] = big.NewInt(0)
	memorize[n].Add(memorize[n], fib(n-1))
	memorize[n].Add(memorize[n], fib(n-2))
	//https://pkg.go.dev/math/big#Int.Add
	return memorize[n]
}
