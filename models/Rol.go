package models

import "gorm.io/gorm"

type Rol struct {
	gorm.Model
	Id      uint64    ` json:"id" gorm:"primary_key;autoIncrement" `
	Nombre  string    `json:"nombre" gorm:"unique;not null"`
	Estado  bool      `json:"estado" gorm:"default:true"`
	Usuario []Usuario `json:"usuarios"`
}

func (Rol) TableName() string {
	return "roles"
}
