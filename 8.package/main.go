package main

import (
	"fmt"
	_ "tang.com/go-study/8.package/hello"
	"tang.com/go-study/8.package/lib"
)

func main() {
	fmt.Println("main init...")
	lib.Test()

	//var slice = make([]int, 1, 1)

}
