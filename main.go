package main

import (
	"log"
	"net/http"

	"github.com/henvisi-1994/api-golang/data"
	"github.com/henvisi-1994/api-golang/models"
	"github.com/henvisi-1994/api-golang/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env", err)
	}
	data.ConectarPostgres()
	data.DB.AutoMigrate(&models.Rol{})
	data.DB.AutoMigrate(&models.Usuario{})
	rutas := routes.InitRoute()
	log.Fatal(http.ListenAndServe(":3001", rutas))
}
