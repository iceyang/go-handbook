package main

import (
	"testing"
)

func TestTypeOf(t *testing.T) {
	t.Logf("type of [...]int{1,2,3,4,5}: %T\n", [...]int{1, 2, 3, 4, 5})
	t.Logf("type of []int{1,2,3,4,5}: %T\n", []int{1, 2, 3, 4, 5})
}

func TestArrayInitialize(t *testing.T) {
	// 声明式
	var arr1 [5]int
	t.Log("arr1:", arr1)

	// 直接赋值
	arr2 := [5]int{}
	t.Log("arr2:", arr2)

	// 赋值并初始化值
	arr3 := [5]int{1, 2, 3, 4, 5}
	t.Log("arr3:", arr3)

	// 赋值，部分初始化
	arr4 := [5]int{1, 2, 3}
	t.Log("arr4:", arr4)

	// 赋值，自动计算数组长度
	arr5 := [...]int{1, 2, 3, 4, 5}
	t.Log("arr5:", arr5)

	// 赋值，通过下标初始化
	arr6 := [5]string{0: "first", 2: "third", 1: "second"}
	t.Log("arr6:", arr6)
}

func TestArrayUsage(t *testing.T) {
	var arr7 [5]int
	// arr7: [0 0 0 0 0]
	arr7[0] = 1
	// arr7: [1 0 0 0 0]
	arr7[4] = 1
	// arr7: [1 0 0 0 1]
	t.Log("arr7:", arr7)
}

func TestSliceInitialize(t *testing.T) {
	// 声明式，这是长度为0的切片
	var slice1 []int
	t.Log("slice1: ", slice1, len(slice1), cap(slice1))

	// 赋值初始化，切片长度为5，容量为5
	slice2 := []int{1, 2, 3, 4, 5}
	t.Log("slice2: ", slice2, len(slice2), cap(slice2))

	// 通过make创建长度为5的切片
	slice3 := make([]int, 5)
	t.Log("slice3: ", slice3, len(slice3), cap(slice3))

	// 通过make创建长度为5，容量为10的切片
	slice4 := make([]int, 5, 10)
	t.Log("slice4: ", slice4, len(slice4), cap(slice4))

	// 通过make创建长度为0，容量为10的切片
	slice5 := make([]int, 0, 10)
	t.Log("slice5: ", slice5, len(slice5), cap(slice5))

	// 赋值，通过下标初始化
	slice6 := []string{0: "first", 2: "third", 1: "second"}
	t.Log("slice6: ", slice6, len(slice6), cap(slice6))
}

func TestSliceUsage(t *testing.T) {
	slice7 := make([]int, 5)
	// slice7: [0 0 0 0 0]
	slice7[1] = 2
	// slice7: [0 2 0 0 0]
	slice7[3] = 3
	// slice7: [0 2 0 3 0]
	t.Log("slice7: ", slice7, len(slice7), cap(slice7))

	slice8 := []int{1, 2, 3, 4, 5}
	// 追加单个元素
	slice8 = append(slice8, 6)
	// 追加多个元素
	slice8 = append(slice8, []int{7, 8}...)
	slice8 = append(slice8, 9, 10)
	t.Log("slice8: ", slice8, len(slice8), cap(slice8))
}

func TestSub(t *testing.T) {
	arr8 := [5]int{1, 2, 3, 4, 5}
	slice9 := []int{1, 2, 3, 4, 5}

	t.Log(arr8[1:])
	t.Log(arr8[:3])
	t.Log(slice9[1:3])
}

func TestLenAndCap(t *testing.T) {
	slice10 := make([]int, 5, 10)
	length := len(slice10)
	capacity := cap(slice10)
	t.Log(slice10, length, capacity)
}

func TestRange(t *testing.T) {
	arr9 := [...]int{1, 2, 3, 4, 5}
	for index, value := range arr9 {
		t.Logf("index: %d, value: %d\n", index, value)
	}
}

func TestOutOfBound(t *testing.T) {
	arr10 := [5]int{1, 2, 3, 4, 5}
	// _ = arr10[6] // 编译不通过

	index := 6
	_ = arr10[index]
}
