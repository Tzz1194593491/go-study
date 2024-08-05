package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Account  string
	Password string
}
