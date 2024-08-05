package main

import "fmt"

func main() {
	//1
	var mapVar1 map[string]string
	mapVar1 = make(map[string]string, 10)
	mapVar1["1"] = "java"
	mapVar1["2"] = "go"
	mapVar1["3"] = "php"
	fmt.Println(mapVar1)
	//2
	mapVar2 := make(map[string]string)
	mapVar2["1"] = "java"
	mapVar2["2"] = "go"
	mapVar2["3"] = "python"
	fmt.Println(mapVar2)
	//3
	var mapVar3 = make(map[string]string, 10)
	mapVar3["1"] = "java"
	mapVar3["2"] = "go"
	mapVar3["3"] = "python"
	fmt.Println(mapVar3)

	//遍历
	for index, value := range mapVar3 {
		fmt.Println(index + " " + value)
	}

	delete(mapVar3, "3")
	fmt.Println("----------------")
	mapVar3["1"] = "gogogo"

	//遍历
	for index, value := range mapVar3 {
		fmt.Println(index + " " + value)
	}

}
