package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jeffveleze/gu_mvc/db"
)

var dbDriver = "mysql"
var dbName = "gu_dev"
var dbUser = "root"
var dbPassword = "Jefferson034"
var dbCreds = dbUser + ":" + dbPassword + "@/" + dbName

type User struct {
	Id          int     `json:id`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	CreatedDate []uint8 `json:"created_date"`
	Password    string  `json:"password"`
	Token       string  `json:token`
}

type UserModel struct {
	dbClient db.DbClient
}

func NewUserModel(db *db.DbClient) *UserModel {
	return &UserModel{*db}
}

func (model UserModel) GetUserByID(userID int) User {

	// db, err := sql.Open(dbDriver, dbCreds)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// stmtOut, err := db.Prepare("SELECT * FROM users WHERE id = ?")

	stmtOut, err := model.dbClient.Database.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	var user User

	err = stmtOut.QueryRow(userID).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedDate, &user.Password, &user.Token)
	if err != nil {
		panic(err.Error())
	}

	return user
}

func (m UserModel) GetAllUsers() []User {
	var users []User

	user1 := User{Id: 1, Name: "Jeff"}
	user2 := User{Id: 2, Name: "Steve"}

	users = append(users, user1)
	users = append(users, user2)

	return users
}
