package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/henvisi-1994/api-golang/data"
	"github.com/henvisi-1994/api-golang/models"
	"github.com/henvisi-1994/api-golang/utils"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	var roles []models.Rol
	data.DB.Find(&roles)
	respuesta := utils.Respuesta{
		Message:    "Listado de roles",
		Data:       roles,
		StatusCode: 200,
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}
func NewRol(w http.ResponseWriter, r *http.Request) {
	var rol models.Rol
	w.Header().Set("content-type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&rol); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Message:    "Error en los datos enviados",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	nuevoRol := data.DB.Create(&rol)

	if nuevoRol.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Message:    "Error al intentar crear Rol",
			StatusCode: http.StatusInternalServerError,
			Data:       nuevoRol.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Message:    "Rol creado Exitosamente",
		StatusCode: http.StatusCreated,
		Data:       rol,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}

func GetRol(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rol models.Rol
	data.DB.First(&rol, params["id"])
	if rol.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		respuesta := utils.Respuesta{
			Message:    "Rol no encontrado",
			Data:       nil,
			StatusCode: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	data.DB.Model(&rol).Association("Usuarios").Find(&rol.Usuarios)
	respuesta := utils.Respuesta{
		Message:    "Rol encontrado",
		Data:       rol,
		StatusCode: http.StatusOK,
	}
	json.NewEncoder(w).Encode(respuesta)

}
func DeleteRol(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rol models.Rol
	data.DB.First(&rol, params["id"])
	if rol.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		respuesta := utils.Respuesta{
			Message:    "Rol no encontrado",
			Data:       nil,
			StatusCode: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	//data.DB.Unscoped().Delete(&rol)
	data.DB.Delete(&rol)
	respuesta := utils.Respuesta{
		Message:    "Rol encontrado",
		Data:       rol,
		StatusCode: http.StatusOK,
	}
	json.NewEncoder(w).Encode(respuesta)

}

func UpdateRol(w http.ResponseWriter, r *http.Request) {
	var rol models.Rol
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	if err := json.NewDecoder(r.Body).Decode(&rol); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Message:    "Error al decodificar solicitud",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	var rol_existente models.Rol

	if err := data.DB.First(&rol_existente, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		respuesta := utils.Respuesta{
			Message:    "Rol no encontrado",
			Data:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	rol_existente.Nombre = rol.Nombre
	rol_existente.Activo = rol.Activo

	if err := data.DB.Save(&rol_existente).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Message:    "Error al intentar actualizar Rol",
			StatusCode: http.StatusInternalServerError,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	respuesta := utils.Respuesta{
		Message:    "Rol actualizado Exitosamente",
		StatusCode: http.StatusCreated,
		Data:       rol,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
