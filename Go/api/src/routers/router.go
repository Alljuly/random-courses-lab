package router

import( "github.com/gorilla/mux"
	"api/src/routers/route"
)


//Retornar as rotas
func Generate() *mux.Router {
	r := mux.NewRouter()

	return route.ConfigRoutes(r)
}