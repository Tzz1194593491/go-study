package main

import (
	"fmt"
	. "polymorphic/interface"
)

type Normal struct {
	Age  int
	Name string
}

func (_this Normal) Breath() {
	fmt.Println(_this, "is Breathing...")
}

func (_this Normal) Eat() {
	fmt.Println(_this, "is Eating")
}

func (_this Normal) GetAge() int {
	return _this.Age
}

func handle(object AnyInterface) {
	value, ok := object.(Normal)
	if ok {
		fmt.Println(value)
	}
}

func main() {
	var person Normal
	person.Name = "Alen"
	person.Age = 18

	var superPerson Person
	superPerson = &person
	superPerson.Eat()
	handle(person)
}
