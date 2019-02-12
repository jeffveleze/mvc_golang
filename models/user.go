package models

import (
	"database/sql"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jeffveleze/gu_mvc/db"
	"github.com/jeffveleze/gu_mvc/entities"
)

type UserModel struct {
	dbClient *db.DbClient
}

func NewUserModel(db *db.DbClient) *UserModel {
	return &UserModel{
		dbClient: db,
	}
}

func (m UserModel) GetUserByID(userID int) entities.User {

	stmtOut, err := m.dbClient.Database.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var user entities.User

	err = stmtOut.QueryRow(userID).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedDate, &user.Password, &user.Token)
	if err != nil {
		panic(err.Error())
	}

	return user
}

func (m UserModel) GetAllUsers() []entities.User {

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

	var users []entities.User

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var user entities.User

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

func (m UserModel) CreateNewUser(user entities.User) entities.User {

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

	var userCreated entities.User

	err = stmtOut.QueryRow().Scan(&userCreated.Id, &userCreated.Name, &userCreated.Email, &userCreated.CreatedDate, &userCreated.Password, &userCreated.Token)
	if err != nil {
		panic(err.Error())
	}

	return userCreated
}

func (m UserModel) DeleteUser(userID int) entities.QueryResult {

	stmtDel, err := m.dbClient.Database.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDel.Close()

	_, err = stmtDel.Exec(userID)
	if err != nil {
		panic(err.Error())
	}

	queryResult := entities.QueryResult{Status: "Query excecuted without errors"}

	return queryResult
}

func (m UserModel) IsAuthorized(user entities.User) (entities.User, error) {

	stmtOut, err := m.dbClient.Database.Prepare("SELECT * FROM users WHERE email = ? AND password = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var dbUser entities.User
	err = stmtOut.QueryRow(user.Email, user.Password).Scan(&dbUser.Id, &dbUser.Name, &dbUser.Email, &dbUser.CreatedDate, &dbUser.Password, &dbUser.Token)

	return dbUser, err
}

func (m UserModel) CreateToken(user entities.User) (entities.JwtToken, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":     user.Name,
		"password": user.Password,
	})

	tokenString, err := token.SignedString([]byte(entities.SecretPassword))

	jwtToken := entities.JwtToken{Token: tokenString}

	return jwtToken, err
}

func (m UserModel) UpdateToken(jwtToken entities.JwtToken, user entities.User) error {

	stmtUpd, err := m.dbClient.Database.Prepare("UPDATE users SET token = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtUpd.Close()

	_, err = stmtUpd.Exec(jwtToken.Token, user.Id)

	return err
}
