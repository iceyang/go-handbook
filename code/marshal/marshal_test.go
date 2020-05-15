package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Name string `json:"name"`
}

type Student struct {
	Person `bson:"inline"`
	Num    int `json:"num"`
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

func TestXMLMarshal(t *testing.T) {
	stu := &Student{
		Person{"Justin"},
		16,
	}
	bytes, _ := xml.Marshal(stu)
	fmt.Printf("student: %s\n", bytes)
}
