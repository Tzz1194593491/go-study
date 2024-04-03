package main

import (
	"fmt"
	"reflect"
)

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100

	var x = new(int)
	*x = 100
	fmt.Println(x)
	fmt.Println(*x)
	fmt.Println(reflect.TypeOf(x))

	var y = make(map[string]int, 10)
	y["测试"] = 100
	fmt.Println(y)
}
