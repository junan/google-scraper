package db

import (
	"fmt"
	_ "google-scraper/models"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/orm"
)

func init() {
	// Register postgres driver
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		fmt.Println("Postgres Driver registration failed: ", err)
	}

	// Register the database
	err = orm.RegisterDataBase("default", "postgres", beego.AppConfig.DefaultString("dbUrl"))
	if err != nil {
		fmt.Println("Postgres connection failed: ", err)
	}

	// Register models
	RegisterModels()

	// Keeping the models and database in sync
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Critical(fmt.Sprintf("Failed to sync models with the database %v", err))
	}
}

func RegisterModels() {
	orm.RegisterModel(new(User))
}
