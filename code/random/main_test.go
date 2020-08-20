package main

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	rules := []Rule{
		Rule{"5元", 75},
		Rule{"30元", 20},
		Rule{"50元", 5},
	}
	result := map[string]int{
		"5元":  0,
		"30元": 0,
		"50元": 0,
	}
	for i := 0; i < 1000000; i++ {
		rule := random(rules)
		result[rule.Name] = result[rule.Name] + 1
	}
	for name, times := range result {
		fmt.Printf("%s: %d\n", name, times)
	}
}
