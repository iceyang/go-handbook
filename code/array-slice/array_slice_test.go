package main

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
	No   int
}

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
	t.Log(arr10)
	// _ = arr10[6] // 编译不通过

	// 打开注释会发生运行时错误
	// index := 6
	// _ = arr10[index]
}

func TestComplexType(t *testing.T) {
	// 初始化了一个二维数组
	arr11 := [5][5]int{}
	// 将第三行第三列赋值为1
	arr11[2][2] = 1
	for _, v := range arr11 {
		t.Log(v)
	}

	slice11 := []*Student{
		&Student{"Justin", 1},
	}
	t.Log("slice11[0]: ", slice11[0])
}

func TestAppendTrap(t *testing.T) {
	slice12 := make([]int, 5, 10)
	slice13 := append(slice12, 6)
	slice13[0] = 1
	// 猜猜这里会输出什么
	fmt.Println("slice12: ", slice12)

	slice14 := make([]int, 5, 10)
	slice15 := append(slice14, 6, 7, 8, 9, 10, 11)
	slice15[0] = 1
	// 猜猜这里会输出什么
	fmt.Println("slice14: ", slice14)

	slice16 := make([]int, 5, 10)
	fmt.Println("slice16: ", slice16)

	slice17 := slice16[:10]
	fmt.Println("slice17: ", slice17)

	// slice18 := slice16[:11] // 这里会产生越界错误
}

func TestAppendTrap2(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	s2 = append(s2, 5, 6, 7)
	t.Log("s1:", s1)
	t.Log("s2:", s2)
}

func TestAppendTrap3(t *testing.T) {
	s1 := make([]int, 3, 6)
	s1[0], s1[1] = 1, 2
	s2 := s1[1:]
	s2[0] = 4
	// t.Logf("len(s2):%d, cap(s2):%d\n", len(s2), cap(s2))
	s2 = append(s2, 6)
	s2[1] = 5
	t.Log("s1:", append(s1, 7))
	t.Log("s2:", s2)
}

/**
 * Slice传入函数中，是传引用
 */
func setSliceValue(arr []Student) {
	arr[0].Name = "Test"
}

func TestSliceReference(t *testing.T) {
	arr := []Student{
		Student{"Justin", 18},
	}
	fmt.Println("(Slice) Before: ", arr[0].Name)
	setSliceValue(arr)
	fmt.Println("(Slice) After: ", arr[0].Name)
}

/**
 * 数组传入函数中，是拷贝
 */
func setArrValue(arr [1]Student) {
	arr[0].Name = "Test"
}

func TestArrReference(t *testing.T) {
	arr := [1]Student{
		Student{"Justin", 18},
	}
	fmt.Println("(Slice) Before: ", arr[0].Name)
	setArrValue(arr)
	fmt.Println("(Slice) After: ", arr[0].Name)
}
