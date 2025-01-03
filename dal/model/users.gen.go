// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        int32          `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Username  string         `gorm:"column:username" json:"username"`
	Password  string         `gorm:"column:password" json:"password"`
	Email     string         `gorm:"column:email" json:"email"`
	Role      string         `gorm:"column:role" json:"role"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
