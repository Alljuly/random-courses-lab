package route

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{

	{
		URI: "/users",
		Method: http.MethodPost,
		Fn: controllers.CreateUser,
		IsAuth: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Fn: controllers.GetUsers,
		IsAuth: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Fn: controllers.GetUserByID,
		IsAuth: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Fn: controllers.UpdateUserByID,
		IsAuth: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Fn: controllers.DeleteUserByID,
		IsAuth: false,
	},
}
