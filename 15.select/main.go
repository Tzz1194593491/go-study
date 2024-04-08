package main

import "fmt"

func main() {
	//select具备多路channel的功能
	//数据管道
	c := make(chan int)
	//监控管道
	quit := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	x, y := 1, 1
	for {
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}
