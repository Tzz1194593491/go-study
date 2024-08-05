package user

type User struct {
	Name string
	Age  int
}

func (this *User) GetName() string {
	return this.Name
}

func (this *User) GetAge() int {
	return this.Age
}
