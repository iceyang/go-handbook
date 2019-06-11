package main

import "testing"

func TestMap(t *testing.T) {
	_ = map[string]string{}

	a := map[string][]int{}
	a["test"] = []int{1, 2, 3}

	studentMap := map[string]Student{}
	studentMap["justin"] = Student{
		Name: "Justin",
		No:   1,
	}

	t.Log(studentMap["justin"])
}
