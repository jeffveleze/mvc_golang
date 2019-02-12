package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jeffveleze/gu_mvc/entities"
	"github.com/jeffveleze/gu_mvc/models"
)

type UserController struct {
	userModel models.UserModel
}

func NewUserController(userModel *models.UserModel) *UserController {
	return &UserController{userModel: *userModel}
}

func (c UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIDString := r.URL.Query().Get("userID")
	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		panic(err)
	}

	user := c.userModel.GetUserByID(userID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func (c UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := c.userModel.GetAllUsers()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

func (c UserController) NewUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	createdUser := c.userModel.CreateNewUser(user)
	jwtToken, err := c.userModel.CreateToken(createdUser)

	if err != nil {
		panic(err)
	}

	err = c.userModel.UpdateToken(jwtToken, createdUser)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		panic(err)
	}
}

func (c UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDString := r.URL.Query().Get("userID")
	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		panic(err)
	}

	queryResult := c.userModel.DeleteUser(userID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(queryResult); err != nil {
		panic(err)
	}
}

func (c UserController) IsAuthorized(w http.ResponseWriter, r *http.Request) {
	userIDString := r.URL.Query().Get("userID")
	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		panic(err)
	}

	queryResult := c.userModel.DeleteUser(userID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(queryResult); err != nil {
		panic(err)
	}
}

func (c UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	authorizedUser, err := c.userModel.IsAuthorized(user)

	if err != nil {
		queryResult := entities.QueryResult{Status: "No authorized user"}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnauthorized)

		if err := json.NewEncoder(w).Encode(queryResult); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(authorizedUser); err != nil {
		panic(err)
	}
}

func (c UserController) CreateToken(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	jwtToken, err := c.userModel.CreateToken(user)
	if err != nil {
		queryResult := entities.QueryResult{Status: "Couldn't create token"}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)

		if err := json.NewEncoder(w).Encode(queryResult); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jwtToken); err != nil {
		panic(err)
	}
}
