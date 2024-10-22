package data

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarPostgres() {
	var error error
	DB, error = gorm.Open(postgres.Open(os.Getenv("CONECTION_STRING")), &gorm.Config{})
	if error != nil {
		log.Fatal("Error de conectexion en la base de datos", error)
	} else {
		log.Println("Conectado a la base de datos postgres")
	}
}
