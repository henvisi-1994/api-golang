package main

import (
	"log"
	"net/http"

	"github.com/henvisi-1994/api-golang/data"
	"github.com/henvisi-1994/api-golang/models"
	"github.com/henvisi-1994/api-golang/routes"
)

func main() {
	data.ConectarPostgres()
	data.DB.AutoMigrate(&models.Rol{})
	data.DB.AutoMigrate(&models.Usuario{})
	rutas := routes.InitRoute()
	log.Fatal(http.ListenAndServe(":3001", rutas))
}
