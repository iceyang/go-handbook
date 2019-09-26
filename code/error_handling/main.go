package main

import (
	"errors"
	"fmt"
)

func division(divisor, dividend int) (int, error) {
	if dividend == 0 {
		return 0, errors.New("dividend cannot be zero")
	}
	return divisor / dividend, nil
}

func main() {
	if res, err := division(10, 5); err != nil {
		fmt.Println(res)
	}
	if res, err := division(10, 0); err != nil {
		fmt.Println(res)
	}
}
