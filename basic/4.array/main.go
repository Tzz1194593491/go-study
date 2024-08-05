package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := '1'
	fmt.Println(p)
	var a = [5]int{1, 2, 3, 4, 5}
	var b [2][2]int
	fmt.Println(a)
	fmt.Println(b)

	c := [2]int{}
	fmt.Printf("c: %p\n", &c)

	test(&c)
	fmt.Println(c)

	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	rand.Seed(time.Now().Unix())

	var d [10]int
	for i := 0; i < 10; i++ {
		// 产生一个0到1000随机数
		d[i] = rand.Intn(1000)
	}
	sum := sumArr(&d)
	fmt.Println(sum)
}

func sumArr(b *[10]int) interface{} {
	var sum int = 0
	for i := 0; i < len(b); i++ {
		sum += b[i]
	}
	return sum
}

func test(x *[2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}
