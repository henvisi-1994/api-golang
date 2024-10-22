package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/henvisi-1994/api-golang/auth"
	"github.com/henvisi-1994/api-golang/data"
	"github.com/henvisi-1994/api-golang/models"
	"github.com/henvisi-1994/api-golang/utils"
)

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuario []models.Usuario
	data.DB.Preload("Rol").Find(&usuario)
	respuesta := utils.Respuesta{
		Message:    "Listado de usuario",
		Data:       usuario,
		StatusCode: 200,
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(respuesta)

}
func NewUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	w.Header().Set("content-type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Message:    "Error en los datos enviados",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	nuevoUsuario := data.DB.Preload("Rol").Create(&usuario)

	if nuevoUsuario.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Message:    "Error al intentar crear Usuario",
			StatusCode: http.StatusInternalServerError,
			Data:       nuevoUsuario.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	if err := data.DB.Preload("Rol").First(&usuario, usuario.ID).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Message:    "Error al intentar crear Usuario",
			StatusCode: http.StatusInternalServerError,
			Data:       nuevoUsuario.Error.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	respuesta := utils.Respuesta{
		Message:    "Usuario creado Exitosamente",
		StatusCode: http.StatusCreated,
		Data:       usuario,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}

func GetUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var usuario models.Usuario
	data.DB.Preload("Rol").First(&usuario, params["id"])
	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		respuesta := utils.Respuesta{
			Message:    "Usuario no encontrado",
			Data:       nil,
			StatusCode: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	respuesta := utils.Respuesta{
		Message:    "Usuario encontrado",
		Data:       usuario,
		StatusCode: http.StatusOK,
	}
	json.NewEncoder(w).Encode(respuesta)

}
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var usuario models.Usuario
	data.DB.First(&usuario, params["id"])
	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		respuesta := utils.Respuesta{
			Message:    "Usuario no encontrado",
			Data:       nil,
			StatusCode: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	//data.DB.Unscoped().Delete(&usuario)
	data.DB.Delete(&usuario)
	respuesta := utils.Respuesta{
		Message:    "Usuario encontrado",
		Data:       usuario,
		StatusCode: http.StatusOK,
	}
	json.NewEncoder(w).Encode(respuesta)

}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respuesta := utils.Respuesta{
			Message:    "Error al decodificar solicitud",
			StatusCode: http.StatusBadRequest,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	var usuario_existente models.Usuario

	if err := data.DB.Preload("Rol").First(&usuario_existente, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		respuesta := utils.Respuesta{
			Message:    "Usuario no encontrado",
			Data:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	usuario_existente.Nombre = usuario.Nombre
	usuario_existente.Correo = usuario.Correo
	usuario_existente.RolId = usuario.RolId

	if err := data.DB.Preload("Rol").Save(&usuario_existente).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respuesta := utils.Respuesta{
			Message:    "Error al intentar actualizar Usuario",
			StatusCode: http.StatusInternalServerError,
			Data:       err.Error(),
		}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	respuesta := utils.Respuesta{
		Message:    "Usuario actualizado Exitosamente",
		StatusCode: http.StatusCreated,
		Data:       usuario,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

type Credenciales struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}
type Claims struct {
	Correo string `json:"correo"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var credenciales Credenciales
	if err := json.NewDecoder(r.Body).Decode(&credenciales); err != nil {
		sendError(w, "Error en los datos de la solicitud", http.StatusBadRequest)
		return
	}

	// Limpiar el correo electrónico
	credenciales.Correo = strings.TrimSpace(credenciales.Correo)

	var usuario models.Usuario
	if err := data.DB.Where("correo = ?", credenciales.Correo).First(&usuario).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			sendError(w, "Datos de acceso incorrectos", http.StatusUnauthorized)
		} else {
			sendError(w, "Error al consultar usuario", http.StatusInternalServerError)
		}
		return
	}

	// Verificación de la contraseña
	if err := VerificarPassword(string(usuario.Password), credenciales.Password); err != nil {
		sendError(w, "Datos de acceso incorrectos", http.StatusUnauthorized)
		return
	}

	// Generar token JWT
	jwtToken, err := auth.GenerarToken(usuario.Correo)
	if err != nil {
		sendError(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	json.NewEncoder(w).Encode(utils.Respuesta{
		Message:    "Autenticación exitosa",
		StatusCode: http.StatusOK,
		Data:       jwtToken,
	})
}

func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}

func VerificarPassword(passwordHashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}
