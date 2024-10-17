package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/henvisi-1994/api-golang/utils"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	respuesta := utils.Respuesta{
		Message:    "Listado de roles",
		Data:       " ",
		StatusCode: 200,
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}
func NewRol(w http.ResponseWriter, r *http.Request) {
	respuesta := utils.Respuesta{
		Message:    "Nuevo Rol",
		Data:       "Nuevo",
		StatusCode: 200,
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}
func GetRol(w http.ResponseWriter, r *http.Request) {
	respuesta := utils.Respuesta{
		Message:    "Buscar",
		Data:       "Buscar uno",
		StatusCode: 200,
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}
func DeleteRol(w http.ResponseWriter, r *http.Request) {
	respuesta := utils.Respuesta{
		Message:    "Borando",
		Data:       "Elimino un rol",
		StatusCode: 200,
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}
