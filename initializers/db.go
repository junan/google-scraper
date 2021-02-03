package initializers

import (
	"fmt"

	"google-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	LoadAppConfig()

	dbUrl, err := web.AppConfig.String("dbUrl")
	if err != nil {
		fmt.Println("Postgres database source is not found: ", err)
	}

	// Register postgres driver
	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		fmt.Println("Postgres driver registration failed: ", err)
	}

	// Register the database
	err = orm.RegisterDataBase("default", "postgres", dbUrl)
	if err != nil {
		fmt.Println("Postgres connection failed: ", err)
	}

	// Register models
	registerModels()

	// Keeping the models and database in sync
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println("Failed to sync models with the database:", err)
	}

	mode, _ := web.AppConfig.String("runmode")
	if mode == "dev" {
		orm.Debug = true
	}
}

func registerModels() {
	orm.RegisterModel(new(models.User))
}
