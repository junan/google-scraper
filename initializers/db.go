package initializers

import (
	"google-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	_ "github.com/beego/beego/v2/server/web/session/redis"
)

func init() {
	LoadAppConfig()

	dbUrl, err := web.AppConfig.String("dbUrl")
	if err != nil {
		logs.Critical("Postgres database source is not found: ", err)
	}

	// Register postgres driver
	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Critical("Postgres driver registration failed: ", err)
	}

	// Register the database
	err = orm.RegisterDataBase("default", "postgres", dbUrl)
	if err != nil {
		logs.Critical("Postgres connection failed: ", err)
	}

	// Register models
	registerModels()

	// Keeping the models and database in sync
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Critical("Failed to sync models with the database:", err)
	}

	mode, _ := web.AppConfig.String("runmode")
	if mode == "dev" {
		orm.Debug = true
	}
}

func registerModels() {
	orm.RegisterModel(new(models.User), new(models.SearchResult))
}
