package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luccascleann/projeto_api/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	rows, err := models.Delete(username)
	if err != nil {
		log.Printf("Erro ao remover o registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"StatusCode": 200,
		"Message":    "usuário apagado",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
