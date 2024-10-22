package auth

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
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
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	//fmt.Println("private key", os.Getenv("API_SECRET_KEY"))
	ecdsaPrivateKey, err := CargarECDSAKeyDesdeEnv()
	if err != nil {
		return "", err
	}

	return jwtToken.SignedString(ecdsaPrivateKey) // Usar la clave privada ECDSA

}
func ValidarToken(r *http.Request) error {
	jwtToken := ExtraerToken(r) // Asegúrate de tener esta función implementada
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método inesperado: %s", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
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
func CargarECDSAKeyDesdeEnv() (*ecdsa.PrivateKey, error) {
	// Leer la clave privada desde la variable de entorno
	ecdsaPrivateKeyPEM := os.Getenv("API_SECRET_KEY")
	if ecdsaPrivateKeyPEM == "" {
		return nil, fmt.Errorf("la variable de entorno API_SECRET_KEY no está establecida")
	}

	// Decodificar el bloque PEM
	block, _ := pem.Decode([]byte(ecdsaPrivateKeyPEM))
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("fallo al decodificar el bloque PEM o tipo incorrecto")
	}

	// Parsear la clave privada en formato DER
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error al parsear la clave privada ECDSA: %v", err)
	}

	return privateKey, nil
}
