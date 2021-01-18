package models

type User struct {
	Base

	Name              string
	Email             string `orm:"unique"`
	HashedPassword string
}

// Beego by default creates the table name as singular, it will make it plural
func (u *User) TableName() string {
	return "users"
}
