package test

type Book struct {
	Title  string
	Author string
}

func (_this *Book) SetTitle(newName string) {
	_this.Title = newName
}
