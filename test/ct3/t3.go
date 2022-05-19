package ct3

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 돈 관련 데이터 변환
func T3() {
	input := "15.93"
	pennies, err := convertStringDollarsToPennies(input)

	if err == nil {
		fmt.Println(pennies)
	}
}

func convertStringDollarsToPennies(amount string) (int64, error) {
	// check if amount can convert to a valid float
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	// split the value on "."
	groups := strings.Split(amount, ".")

	// if there is no . result will still be
	// captured here
	result := groups[0]

	// base string
	r := ""

	// handle the data after the "."
	if len(groups) == 2 {
		if len(groups[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = groups[1]
	}

	// pad with 0, this will be
	// 2 0's if there was no .
	for len(r) < 2 {
		r += "0"
	}

	result += r

	// convert it to an int
	return strconv.ParseInt(result, 10, 64)
}
