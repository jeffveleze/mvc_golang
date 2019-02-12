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

func buildRoutes(userController *controllers.UserController) *Routes {
	return &Routes{
		Route{
			"home",
			"GET",
			"/users",
			validateMiddleware(userController.GetUserByID),
		},
		Route{
			"home",
			"GET",
			"/users/all",
			validateMiddleware(userController.GetAllUsers),
		},
		Route{
			"newUser",
			"POST",
			"/users/new",
			userController.NewUser,
		},
		Route{
			"deleteUser",
			"DELETE",
			"/users",
			validateMiddleware(userController.DeleteUser),
		},
		Route{
			"login",
			"POST",
			"/users/login",
			validateMiddleware(userController.Login),
		},
		Route{
			"createToken",
			"POST",
			"/users/create-token",
			userController.CreateToken,
		},
	}
}

func NewRouter(dbClient *db.DbClient) *mux.Router {

	routes := buildRoutes(controllers.NewUserController(models.NewUserModel(dbClient)))

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range *routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
