package route

import ("net/http"
 	"github.com/gorilla/mux"
)

//struct de rotas da api
type Route struct {
	URI    string
	Method string
	Fn     func(http.ResponseWriter, *http.Request)
	IsAuth bool
}

//Insere rotas de usuario em router
func ConfigRoutes(r *mux.Router) *mux.Router{
	routes := usersRoutes

	for _, route := range routes{
		r.HandleFunc(route.URI, route.Fn).Methods(route.Method)
	}

	return r
}