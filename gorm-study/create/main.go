package main

import (
	"fmt"
	"github.com/Tzz1194593491/gorm-study/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = db.AutoMigrate(&model.User{})

	createSingle(db)
	createMany(db)

	createBySomeFields(db)
}

// 创建数据

func createSingle(db *gorm.DB) {
	user := &model.User{
		Account:  "123",
		Password: "123",
	}
	result := db.Create(user)
	fmt.Printf("%+v", result)
}

func createMany(db *gorm.DB) {
	users := []*model.User{
		{Account: "1", Password: "1"},
		{Account: "2", Password: "2"},
		{Account: "3", Password: "3"},
	}
	result := db.Create(users)
	fmt.Printf("%+v", result)
}

// 创建指定列的数据

func createBySomeFields(db *gorm.DB) {
	user := &model.User{
		Account:  "123",
		Password: "123",
	}
	db.Select("account").Create(user)
	fmt.Printf("%+v", user)
}
