package models

import (
	"time"
)

type User struct {
	Id                int
	Name              string
	Email             string `orm:"unique"`
	EncryptedPassword string
	CreatedAt         time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt         time.Time `orm:"auto_now;type(datetime)"`
}

// Beego by default creates the table name as singular, it will make it plural
func (u *User) TableName() string {
	return "users"
}
