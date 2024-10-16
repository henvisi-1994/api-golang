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
	return rutas
}