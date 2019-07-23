package main

import (
	"fmt"
	"testing"
)

func TestStringExample(t *testing.T) {
	_ = "I am a string"

	str := "我是中文字符串"

	t.Logf("Length of '%s': %d\n", str, len(str))
	t.Logf("Length of '%s'(convert to []rune): %d\n", str, len([]rune(str)))

	t.Log(str[:3])
	t.Log(string([]rune(str)[1:]))

	s := "Go编程"
	fmt.Println(len(s))
}
