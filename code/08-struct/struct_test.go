package main

import (
	"fmt"
	"testing"
)

func TestStruct1(t *testing.T) {
	// 定义初始化，数值为默认值
	var person Person
	fmt.Println(person)

	// 通过字段名称初始化
	person1 := Person{
		name: "Justin",
		age:  18,
	}
	fmt.Println(person1.GetName())

	// 按字段顺序初始化
	person2 := Person{"Justin", 18}
	fmt.Println(person2)

	// 部分初始化
	person3 := Person{
		name: "Justin",
	}
	fmt.Println(person3)

	// 使用 new() 函数初始化
	person4 := new(Person)
	person4.name = "Tim"
	fmt.Println(person4)
}

/**
 * 执行结果，打印出来的是：
 * name:
 * SetName:  Justin
 * SetName2:  Justin
 */
func TestStruct2(t *testing.T) {
	person := Person{}
	fmt.Println("name:", person.GetName())
	person.SetName("Justin")
	fmt.Println("SetName: ", person.GetName())
	person.SetName2("Justin2")
	fmt.Println("SetName2: ", person.GetName())
}

func TestStruct3(t *testing.T) {
	o := AnonymousField{}
	o.int = 10
	fmt.Println(o)
}

func TestStruct4(t *testing.T) {
	adult := &Adult{}
	adult.job = "salesman"
	adult.person.name = "John"
	adult.person.age = 25
	fmt.Println(adult)
}
