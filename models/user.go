package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id                int
	Name              string
	Email             string `orm:"unique"`
	EncryptedPassword string
	CreatedAt         time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt         time.Time `orm:"auto_now;type(datetime); null"`
}

func init() {
	orm.RegisterModel(new(User))
}
