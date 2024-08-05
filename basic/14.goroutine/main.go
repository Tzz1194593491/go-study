package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		fmt.Println(i)
		i++
		time.Sleep(time.Second * time.Duration(i))
	}
}

func main() {
	//go newTask()

	go func(a int, b int) {
		fmt.Println(a, b)
	}(1, 2)
	fmt.Println(1)
	//for {
	//	time.Sleep(time.Duration(1) * time.Second)
	//}

}
