package models

import (
	"database/sql"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jeffveleze/gu_mvc/db"
)

var dbDriver = "mysql"
var dbName = "gu_dev"
var dbUser = "root"
var dbPassword = "Jefferson034"
var dbCreds = dbUser + ":" + dbPassword + "@/" + dbName
var secretPassword = "secret"

type User struct {
	Id          int     `json:id`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	CreatedDate []uint8 `json:"created_date"`
	Password    string  `json:"password"`
	Token       string  `json:token`
}

type UserModel struct {
	dbClient *db.DbClient
}

func NewUserModel(db *db.DbClient) *UserModel {
	return &UserModel{
		dbClient: db,
	}
}

func (m UserModel) GetUserByID(userID int) User {

	stmtOut, err := m.dbClient.Database.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var user User

	err = stmtOut.QueryRow(userID).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedDate, &user.Password, &user.Token)
	if err != nil {
		panic(err.Error())
	}

	return user
}

func (m UserModel) GetAllUsers() []User {

	rows, err := m.dbClient.Database.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var users []User

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var user User

		user.Id, err = strconv.Atoi(string(values[0]))
		user.Name = string(values[1])
		user.Email = string(values[2])
		user.CreatedDate = []uint8(string(values[3]))
		user.Password = string(values[4])
		user.Token = string(values[5])

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return users
}

func (m UserModel) CreateNewUser(user User) User {

	stmtIns, err := m.dbClient.Database.Prepare("INSERT INTO users (name, email, password, token) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	stmtOut, err := m.dbClient.Database.Prepare("SELECT * FROM users ORDER BY id DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	_, err = stmtIns.Exec(user.Name, user.Email, user.Password, "")
	if err != nil {
		panic(err.Error())
	}

	var userCreated User

	err = stmtOut.QueryRow().Scan(&userCreated.Id, &userCreated.Name, &userCreated.Email, &userCreated.CreatedDate, &userCreated.Password, &userCreated.Token)
	if err != nil {
		panic(err.Error())
	}

	return userCreated
}

func (m UserModel) DeleteUser(userID int) QueryResult {

	stmtDel, err := m.dbClient.Database.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDel.Close()

	_, err = stmtDel.Exec(userID)
	if err != nil {
		panic(err.Error())
	}

	queryResult := QueryResult{Status: "Query excecuted without errors"}

	return queryResult
}

func (m UserModel) IsAuthorized(user User) (User, error) {

	stmtOut, err := m.dbClient.Database.Prepare("SELECT * FROM users WHERE email = ? AND password = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var dbUser User
	err = stmtOut.QueryRow(user.Email, user.Password).Scan(&dbUser.Id, &dbUser.Name, &dbUser.Email, &dbUser.CreatedDate, &dbUser.Password, &dbUser.Token)

	return dbUser, err
}

func (m UserModel) CreateToken(user User) (JwtToken, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":     user.Name,
		"password": user.Password,
	})

	tokenString, err := token.SignedString([]byte(secretPassword))

	jwtToken := JwtToken{Token: tokenString}

	return jwtToken, err
}

func (m UserModel) UpdateToken(jwtToken JwtToken, user User) error {

	stmtUpd, err := m.dbClient.Database.Prepare("UPDATE users SET token = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtUpd.Close()

	_, err = stmtUpd.Exec(jwtToken.Token, user.Id)

	return err
}
