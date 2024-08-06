package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Printf("BeforeCreate-%+v", u)
	return nil
}
