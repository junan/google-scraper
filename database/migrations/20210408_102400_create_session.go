package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
	"github.com/beego/beego/v2/core/logs"
)

// DO NOT MODIFY
type CreateSession_20210408_102400 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateSession_20210408_102400{}
	m.Created = "20210408_102400"

	err := migration.Register("CreateSession_20210408_102400", m)
	if err != nil {
		logs.Error("Migration failed: ", err)
	}
}

// Run the migrations
func (m *CreateSession_20210408_102400) Up() {
	m.SQL(`CREATE TABLE "session"
		(
			session_key	char(64) NOT NULL,
			session_data	bytea,
			session_expiry	timestamp NOT NULL,
			CONSTRAINT session_key PRIMARY KEY(session_key)
		);`)
}

// Reverse the migrations
func (m *CreateSession_20210408_102400) Down() {
	m.SQL(`DROP TABLE "session";`)
}
