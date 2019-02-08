package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffveleze/gu_mvc/controllers"
	"github.com/jeffveleze/gu_mvc/db"
	"github.com/jeffveleze/gu_mvc/models"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var userModel models.UserModel
var userController controllers.UserController

var routes = Routes{
	Route{
		"home",
		"GET",
		"/users",
		userController.GetUserByID,
	},
	Route{
		"home",
		"GET",
		"/users/all",
		userController.GetAllUsers,
	},
}

func NewRouter(dbClient *db.DbClient) *mux.Router {

	userModel = *models.NewUserModel(dbClient)
	userController = *controllers.NewUserController(&userModel)

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
