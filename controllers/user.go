package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jeffveleze/gu_mvc/models"
)

type UserController struct {
	userModel models.UserModel
}

func NewUserController(userModel models.UserModel) UserController {
	return UserController{userModel: userModel}
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
	user := c.userModel.GetAllUsers()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}
