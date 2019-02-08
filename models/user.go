package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jeffveleze/gu_mvc/db"
)

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
