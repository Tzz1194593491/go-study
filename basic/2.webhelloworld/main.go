package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Hello, World!")
		if err != nil {
			return
		}
	})
	fmt.Println("Please Visit -  http://localhost:8888/")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}
}
