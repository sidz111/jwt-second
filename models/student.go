package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Class    string `json:"class"`
	Token    string `json:"token"`
	Password string `json:"password"`
}
