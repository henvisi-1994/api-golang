package data

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var CONECTION_STRING = "host=172.17.0.3 user=postgres password=ligaHer2025* dbname=api_golang port=5432 sslmode=disable TimeZone=America/Bogota"
var DB *gorm.DB

func ConectarPostgres() {
	var err error
	DB, err = gorm.Open(postgres.Open(CONECTION_STRING), &gorm.Config{})
	if err != nil {
		log.Fatal("Error de conectexion en la base de datos", err)
	} else {
		log.Println("Conectado a la base de datos postgres")
	}
}
