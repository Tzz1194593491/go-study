package main

import "fmt"

func main() {
	/**
	创建切片的方式
	*/
	//1
	var s1 []int
	if s1 == nil {
		fmt.Println("为空")
	} else {
		fmt.Println("不为空")
	}
	//2
	s2 := []int{}
	//3
	var s3 = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)

	var s7 []int = make([]int, 10)
	s7 = make([]int, 10)
	fmt.Println(s7)
	s7 = make([]int, 10, 20)

	var s8 = make([]int, 10, 10)
	fmt.Printf("%p\n", &s8)
	fmt.Println(s8)
	s8 = append(s8, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("%p\n", &s8)
	fmt.Println(s8)
	s9 := s8[:4]
	s9[3] = 1
	fmt.Println(s8)
	fmt.Println(s9)

	s9 = append(s9, 1)
	fmt.Println(s9)
	fmt.Println(s8)

	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d,%p\n", c, n, &c)
			c = n
		}
	}

	slice := make([]int, 4, 4)
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))

	array := []int{10, 20, 30, 40}
	sliceOther := make([]int, 6)
	n := copy(sliceOther, array)
	fmt.Println(n, sliceOther, array)
}
