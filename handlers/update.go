package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luccascleann/projeto_api/models"
)

func Update(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, "username inválido", http.StatusBadRequest)
		return
	}

	rows, err := models.Update(username, user)
	if err != nil {
		log.Printf("Erro ao atualizar o registro: %v", err)
		http.Error(w, "usuário não encontrado", http.StatusNotFound)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"StatusCode": 200,
		"Message":    "usuário atualizado",
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
