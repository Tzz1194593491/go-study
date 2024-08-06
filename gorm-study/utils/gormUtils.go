package utils

import (
	"fmt"
	"github.com/Tzz1194593491/gorm-study/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Info),
		CreateBatchSize: 2,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	_ = db.AutoMigrate(&model.User{})
	return db, nil
}
