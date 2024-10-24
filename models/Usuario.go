package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre   string `json:"nombre" gorm:"size:100;not null"`
	Correo   string `json:"correo" gorm:"size:100;unique;not null"`
	Password string `json:"password" gorm:"not null"`
	RolId    uint   ` json:"rol_id" `
	Rol      Rol    `json:"rol"`
}

func (Usuario) TableName() string {
	return "usuarios"
}
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
func (u *Usuario) BeforeSave(tx *gorm.DB) error {
	passwordHashed, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(passwordHashed)
	return nil
}
func (u *Usuario) Prepare() {
	u.ID = 0
	u.Nombre = html.EscapeString(strings.ToUpper(strings.TrimSpace(u.Nombre)))
	u.Correo = html.EscapeString(strings.TrimSpace(u.Correo))
}
