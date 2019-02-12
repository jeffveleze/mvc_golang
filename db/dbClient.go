package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jeffveleze/gu_mvc/entities"
)

type DbClient struct {
	Database *sql.DB
}

func NewDbClient() *DbClient {
	db, err := sql.Open(entities.DBDriver, entities.DBCreds)
	if err != nil {
		panic(err.Error())
	}

	return &DbClient{db}
}
