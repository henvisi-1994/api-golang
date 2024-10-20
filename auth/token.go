package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerarToken(correo string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = correo
	claims["authorize"] = true
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return jwtToken.SignedString([]byte(os.Getenv("API_SECRET_KEY")))

}
