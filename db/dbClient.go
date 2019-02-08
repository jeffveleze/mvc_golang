package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbDriver = "mysql"
var dbName = "gu_dev"
var dbUser = "root"
var dbPassword = "Jefferson034"
var dbCreds = dbUser + ":" + dbPassword + "@/" + dbName

type DbClient struct {
	Database *sql.DB
}

func NewDbClient() *DbClient {
	db, err := sql.Open(dbDriver, dbCreds)
	if err != nil {
		panic(err.Error())
	}

	return &DbClient{db}
}
