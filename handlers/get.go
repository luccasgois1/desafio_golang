package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luccascleann/projeto_api/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if isUserNameValid, err := ValidateUserName(w, username); !isUserNameValid || err != nil {
		log.Printf("Erro ao validar o username: %v", err)
		return
	}

	user, err := models.Get(username)
	if err != nil {
		log.Printf("Erro ao atualizar o registro: %v", err)
		if err == sql.ErrNoRows {
			http.Error(w, "404 - Usuário não encontrado", http.StatusNotFound)
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - operação realizada com sucesso\n"))
	json.NewEncoder(w).Encode(user)
}
