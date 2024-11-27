package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Role     string
}

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserId  uint
	User    User `gorm:"foreignKey:UserId"`
}
