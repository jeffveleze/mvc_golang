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
	database *sql.DB
}

type DbUser struct {
	Id          int     `json:id`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	CreatedDate []uint8 `json:"created_date"`
	Password    string  `json:"password"`
	Token       string  `json:token`
}

func NewDbClient() DbClient {
	db, err := sql.Open(dbDriver, dbCreds)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// db, err := sql.Open(dbDriver, dbCreds)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()

	return DbClient{db}
}

func (client DbClient) GetUserByID(userID int) string {

	_, err := client.database.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	// stmtOut, err := client.database.Prepare("SELECT * FROM users WHERE id = ?")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer stmtOut.Close()

	// var user DbUser

	// err = stmtOut.QueryRow(userID).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedDate, &user.Password, &user.Token)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// return user.Email

	return "nombre"
}
