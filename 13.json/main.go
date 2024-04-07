package main

import (
	"encoding/json"
	"fmt"
	. "study/13.json/pojo"
)

func main() {
	user := User{Name: "xiaomi", Age: 18}
	userStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", userStr)
	var newUser User
	err = json.Unmarshal(userStr, &newUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newUser)
}
