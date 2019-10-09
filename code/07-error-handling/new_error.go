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

func demo1() {
	if res, err := division(10, 5); err == nil {
		fmt.Println(res)
	}
	res, err := division(10, 0)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
