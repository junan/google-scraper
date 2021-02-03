package test_helpers

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func TruncateTable(tableName string) {
	o := orm.NewOrm()
	rqwSql := fmt.Sprintf("TRUNCATE TABLE \"%s\";", tableName)
	_, err := o.Raw(rqwSql).Exec()
	if err != nil {
		err := orm.RunSyncdb("default", false, false)
		logs.Error("Failed to truncate table", err)
	}
}
