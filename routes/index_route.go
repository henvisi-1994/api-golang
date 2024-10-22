package routes

import (
	"github.com/gorilla/mux"
	"github.com/henvisi-1994/api-golang/controllers"
	"github.com/henvisi-1994/api-golang/middleware"
)

func InitRoute() *mux.Router {
	// Inicializar la conexión con la base de datos
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", controllers.InitRoute).Methods("GET")
	api_roles := api.PathPrefix("/roles").Subrouter()
	api_roles.HandleFunc("", middleware.SetMiddlewareAutentication(controllers.GetRoles)).Methods("GET")
	api_roles.HandleFunc("/{id}", middleware.SetMiddlewareAutentication(controllers.GetRol)).Methods("GET")
	api_roles.HandleFunc("/{id}", middleware.SetMiddlewareAutentication(controllers.UpdateRol)).Methods("PUT")
	api_roles.HandleFunc("", middleware.SetMiddlewareAutentication(controllers.NewRol)).Methods("POST")
	api_roles.HandleFunc("/{id}", middleware.SetMiddlewareAutentication(controllers.DeleteRol)).Methods("DELETE")
	api_usuarios := api.PathPrefix("/usuarios").Subrouter()
	api_usuarios.HandleFunc("", middleware.SetMiddlewareAutentication(controllers.GetUsuarios)).Methods("GET")
	api_usuarios.HandleFunc("/{id}", middleware.SetMiddlewareAutentication(controllers.GetUsuario)).Methods("GET")
	api_usuarios.HandleFunc("/{id}", middleware.SetMiddlewareAutentication(controllers.UpdateUsuario)).Methods("PUT")
	api_usuarios.HandleFunc("", middleware.SetMiddlewareAutentication(controllers.NewUsuario)).Methods("POST")
	api_usuarios.HandleFunc("/{id}", middleware.SetMiddlewareAutentication(controllers.DeleteUsuario)).Methods("DELETE")
	api_auth := api.PathPrefix("/auth").Subrouter()
	api_auth.HandleFunc("/login", controllers.Login).Methods("POST")
	api_auth.HandleFunc("/register", controllers.NewUsuario).Methods("POST")

	return rutas
}
