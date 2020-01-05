package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Name string
}

type Student struct {
	Person `bson:"inline"`
	Num    int
}

func TestJsonMarshal(t *testing.T) {
	stu := &Student{
		Person{"Justin"},
		16,
	}
	bytes, _ := json.Marshal(stu)
	fmt.Printf("student: %s\n", bytes)
}

func TestBsonMarshal(t *testing.T) {
	stu := &Student{
		Person{"Justin"},
		16,
	}
	bytes, _ := bson.Marshal(stu)
	fmt.Printf("student: %s\n", bytes)
	var stu2 Student
	bson.Unmarshal(bytes, &stu2)
	fmt.Println(stu2)
}
