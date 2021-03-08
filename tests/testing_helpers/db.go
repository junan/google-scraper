package testing_helpers

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func TruncateTables(tableNames ...string) {
	var rawSql string
	o := orm.NewOrm()

	for _, tableName := range tableNames {
		rawSql += fmt.Sprintf("TRUNCATE TABLE \"%s\";", tableName)
	}

	_, err := o.Raw(rawSql).Exec()
	if err != nil {
		err := orm.RunSyncdb("default", false, false)
		logs.Error("Failed to truncate table", err)
	}
}

