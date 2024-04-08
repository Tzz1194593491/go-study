package main

import (
	"fmt"
	"time"
)

func main() {
	//channelFirst()
	//channelSecond()
	channelThird()
}

func channelThird() {
	c := make(chan int, 10)

	go func() {
		defer close(c)
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for data := range c {
		fmt.Println(data)
	}
}

func channelSecond() {
	c := make(chan int, 3)
	go func() {
		defer fmt.Println("子go程结束了")
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()

	for i := 0; i < 5; i++ {
		n := <-c
		fmt.Println(n)
	}
	time.Sleep(time.Duration(1) * time.Second)
}

func channelFirst() {
	c := make(chan []byte)

	go read("1", c)
	go read("2", c)
	go read("3", c)

	write(c)
}

func write(c chan []byte) {
	for true {
		var data string
		_, err := fmt.Scan(&data)
		if err != nil {
			fmt.Println(err)
			return
		}
		if data != "" {
			c <- []byte(data)
		}
	}
}

func read(name string, c chan []byte) {
	for {
		s := <-c
		fmt.Println(name + "-" + string(s))
	}
}
