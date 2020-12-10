package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	id	int64 `json:"ID"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	Address []Address `gorm:"ForeignKey:UserID" json:"Address"`
}

type Address struct {
	gorm.Model
	City string `json:"city`
	UserID uint `json:"user_id"`
}
