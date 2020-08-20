package main

import (
	"math/rand"
	"time"
)

type Rule struct {
	Name   string
	Weight int
}

func random(rules []Rule) Rule {
	totalWeight := 0
	for _, rule := range rules {
		totalWeight += rule.Weight
	}
	rand.Seed(int64(time.Now().UnixNano()))
	num := rand.Intn(totalWeight)
	index, count := 0, rules[0].Weight
	for num-count >= 0 {
		index += 1
		count += rules[index].Weight
	}
	return rules[index]
}
