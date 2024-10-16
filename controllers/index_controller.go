package controllers

import (
	"encoding/json"
	"net/http"
)

type Saludo struct {
	Msg        string `json:"message"`
	StatusCode int    `json:"code"`
}

func InitRoute(w http.ResponseWriter, r *http.Request) {
	saludo := Saludo{
		Msg:        "API RESTful con Gorilla Mux",
		StatusCode: 200,
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(saludo)
}
