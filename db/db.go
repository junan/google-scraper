package db

import (
	"fmt"

	"google-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	_ "github.com/lib/pq"
)

func init() {
	dbUrl, _ := web.AppConfig.String("dbUrl")

	// Register postgres driver
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		fmt.Println("Postgres Driver registration failed: ", err)
	}

	// Register the database
	err = orm.RegisterDataBase("default", "postgres", dbUrl)
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

	mode, _ := web.AppConfig.String("runmode")
	if mode == "dev" {
		orm.Debug = true
	}
}

func RegisterModels() {
	orm.RegisterModel(new(models.User))
}
