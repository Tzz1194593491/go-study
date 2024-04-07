package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tty, err := os.OpenFile("D:\\Code\\Golang\\src\\go-study\\13.reflect\\test.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	var r io.Reader
	r = tty
	data := make([]byte, 1024)
	newData := make([]byte, 1024)
	for n, err := r.Read(data); n != 0; n, err = r.Read(data) {
		if err != nil {
			fmt.Println(err)
		}
		newData = append(newData, data...)
	}
	newString := string(newData)
	fmt.Println(newString)

	var w io.Writer
	w = r.(io.Writer)
	_, _ = w.Write([]byte(newString))
}
