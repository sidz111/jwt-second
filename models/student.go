package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"` // this is for authentication and should be unique
	Age      int    `json:"age"`
	Class    string `json:"class"`
	Token    string `json:"token"`
	Password string `json:"password"`
}
