package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	//. "study/12.reflect/user"
)

type User struct {
	Name string
	Age  int
}

func (_this *User) GetName() string {
	return _this.Name
}

func (_this *User) GetAge() int {
	return _this.Age
}

func main() {
	tty, err := os.OpenFile("12.reflect/test.txt", os.O_RDWR, 0)
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

	//var w io.Writer
	//w = r.(io.Writer)
	//_, _ = w.Write([]byte(newString))

	userObj := User{Name: "小明", Age: 18}
	handle(userObj)
	test(&userObj)
}

func handle(inter interface{}) {
	typeValue := reflect.TypeOf(&inter)
	fmt.Println("Type is " + typeValue.Name())

	//value := reflect.ValueOf(inter)
	//for i := 0; i < value.NumField(); i++ {
	//	fmt.Println(typeValue.Field(i).Name, typeValue.Field(i).Type, value.Field(i))
	//}

	for i := 0; i < typeValue.NumMethod(); i++ {
		m := typeValue.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func test(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType)

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	// 如果input是指针，获取其指向的元素的Type
	if inputType.Kind() == reflect.Ptr {
		inputType = inputType.Elem()
		inputValue = inputValue.Elem()
	}

	// 获取字段（Field）
	for i := 0; i < inputType.NumField(); i++ {
		fieldType := inputType.Field(i)
		fieldValue := inputValue.Field(i).Interface()
		fmt.Printf("Field Name: '%s', Field Type: '%s', Field Value: '%v'\n", fieldType.Name, fieldType.Type, fieldValue)
	}

	// 获取方法（Method）
	// 如果需要操作指针接收者的方法，需要再次获取原始指针类型的方法集
	// 由于reflect.TypeOf的值接收者和指针接收者的方法集是不同的
	// 你需要处理指针类型才能拿到全部方法
	if reflect.TypeOf(input).Kind() == reflect.Ptr {
		methodCount := reflect.TypeOf(input).NumMethod()
		for i := 0; i < methodCount; i++ {
			m := reflect.TypeOf(input).Method(i)
			fmt.Printf("Method Name: '%s', Type: '%s'\n", m.Name, m.Type)
		}
	}
}
