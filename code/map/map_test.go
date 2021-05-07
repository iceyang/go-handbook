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

	map3 := map[int]string{
		1: "one",
	}
	t.Log(map3)

	map4 := map[string]Student{}
	v, _ := map4["name"]
	t.Log(v)

	map5 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	delete(map5, "b")

	for k, v := range map5 {
		t.Log(k, v)
	}
}
