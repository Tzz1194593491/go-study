package main

import "fmt"

func fun1() {
	fmt.Println("A")
}

func fun2() {
	fmt.Println("B")
}

func fun3() {
	fmt.Println("C")
}

func deferFun() int {
	fmt.Println("defer")
	return 0
}

func returnFun() int {
	fmt.Println("return")
	return 0
}

func returnAndDefer() int {
	defer deferFun()
	return returnFun()
}

func main() {
	defer fun1()
	defer fun2()
	defer fun3()

	fmt.Println("main...")

	returnAndDefer()
}
