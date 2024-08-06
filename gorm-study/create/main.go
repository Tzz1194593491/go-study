package main

import (
	"fmt"
	"github.com/Tzz1194593491/gorm-study/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Info),
		CreateBatchSize: 2,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = db.AutoMigrate(&model.User{})

	createSingle(db)
	createMany(db)
	createBySomeFields(db)
	createUseHook(db)
	createByMap(db)
}

// 创建数据

func createSingle(db *gorm.DB) {
	user := &model.User{
		Account:  "123",
		Password: "123",
	}
	result := db.Create(user)
	fmt.Printf("%+v\n", result)
}

// 批量插入

func createMany(db *gorm.DB) {
	users := []*model.User{
		{Account: "1", Password: "1"},
		{Account: "2", Password: "2"},
		{Account: "3", Password: "3"},
		{Account: "4", Password: "4"},
	}
	db = db.Session(&gorm.Session{CreateBatchSize: 4})
	result := db.Create(&users)
	for _, user := range users {
		fmt.Printf("%d ", user.ID)
	}
	fmt.Printf("%+v\n", result)
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

// 创建数据时使用钩子

func createUseHook(db *gorm.DB) {
	users := []*model.User{
		{Account: "1", Password: "1"},
		{Account: "2", Password: "2"},
		{Account: "3", Password: "3"},
		{Account: "4", Password: "4"},
	}
	fmt.Printf("createUseHook-skipHooks")
	db.Session(&gorm.Session{SkipHooks: true}).Create(&users)
}

// 创建数据通过map

func createByMap(db *gorm.DB) {
	db.Model(&model.User{}).Create(map[string]interface{}{
		"Account": "jinzhu", "Password": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&model.User{}).Create([]map[string]interface{}{
		{"Account": "jinzhu_1", "Password": "18"},
		{"Account": "jinzhu_2", "Password": "20"},
	})
}
