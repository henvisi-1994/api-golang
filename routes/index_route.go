package routes

import (
	"github.com/gorilla/mux"
	"github.com/henvisi-1994/api-golang/controllers"
)

func InitRoute() *mux.Router {
	// Inicializar la conexión con la base de datos
	rutas := mux.NewRouter()
	api := rutas.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", controllers.InitRoute).Methods("GET")
	api_roles := api.PathPrefix("/roles").Subrouter()
	api_roles.HandleFunc("", controllers.GetRoles).Methods("GET")
	api_roles.HandleFunc("/{id}", controllers.GetRol).Methods("GET")
	api_roles.HandleFunc("/{id}", controllers.UpdateRol).Methods("PUT")
	api_roles.HandleFunc("", controllers.NewRol).Methods("POST")
	api_roles.HandleFunc("/{id}", controllers.DeleteRol).Methods("DELETE")

	return rutas
}
