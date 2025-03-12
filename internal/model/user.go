package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"unique" json:"email"`
}
