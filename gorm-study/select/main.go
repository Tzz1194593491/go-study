package main

import (
	"fmt"
	"github.com/Tzz1194593491/gorm-study/model"
	"github.com/Tzz1194593491/gorm-study/utils"
	"gorm.io/gorm"
)

func main() {
	db, err := utils.CreateDB()
	if err != nil {
		return
	}
	selectSingle(db)
}

func selectSingle(db *gorm.DB) (user *model.User) {
	db.First(&user)
	fmt.Printf("%+v", user)
	db.Take(&user)
	fmt.Printf("%+v", user)
	db.Last(&user)
	fmt.Printf("%+v", user)
	return user
}
