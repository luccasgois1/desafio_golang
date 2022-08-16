package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luccascleann/projeto_api/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	user, err := models.Get(username)
	if err != nil {
		log.Printf("Erro ao atualizar o registro: %v", err)
		http.Error(w, "400 - Usuário não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
