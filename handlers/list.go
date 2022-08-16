package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/luccascleann/projeto_api/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
