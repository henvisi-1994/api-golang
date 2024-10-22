package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerarToken(correo string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = correo
	claims["authorize"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Usar el algoritmo de firma HS256 con una frase secreta
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Cargar la frase secreta desde las variables de entorno
	secret := os.Getenv("API_SECRET_KEY")
	if secret == "" {
		return "", fmt.Errorf("frase secreta no encontrada")
	}

	// Firmar el token con la frase secreta
	return jwtToken.SignedString([]byte(secret))
}
func ValidarToken(r *http.Request) error {
	jwtToken := ExtraerToken(r) // Asegúrate de tener esta función implementada

	// Parsear el token y validar la firma con la frase secreta
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el algoritmo sea HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método inesperado: %s", token.Header["alg"])
		}

		// Cargar la frase secreta desde las variables de entorno
		secret := os.Getenv("API_SECRET_KEY")
		if secret == "" {
			return nil, fmt.Errorf("frase secreta no encontrada")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return fmt.Errorf("error en la validación del token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims) // Procesar los claims si es necesario
		return nil
	}

	return fmt.Errorf("token inválido")
}

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(b))
}
func ExtraerToken(r *http.Request) string {
	parametros := r.URL.Query()
	token := parametros.Get("token")

	if token != "" {
		return token
	}

	tokenString := r.Header.Get("Authorization")
	if len(strings.Split(tokenString, " ")) == 2 {
		return strings.Split(tokenString, " ")[1]
	}
	return ""
}
