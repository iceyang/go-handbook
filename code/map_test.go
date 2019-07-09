package main

import "testing"

type Student struct {
	Name string
	No   int
}

func TestMap(t *testing.T) {
	var map1 map[string]int
	t.Log(map1)

	map2 := map[string]string{}
	map2["name"] = "Bob"
	t.Log(map2)

	val, ok := map2["name2"]
	t.Log(ok)
	t.Log(val)

	map3 := map[string]Student{}
	v, _ := map3["name"]
	t.Log(v)

	map4 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range map4 {
		t.Log(k, v)
	}
}
