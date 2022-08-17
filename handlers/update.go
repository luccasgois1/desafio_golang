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
	//Validate username
	if isUserNameValid, err := ValidateUserName(w, username); !isUserNameValid || err != nil {
		log.Printf("Erro ao validar o username: %v", err)
		return
	}
	// Validate body
	if isBodyValid, err := ValidateBody(w, r); !isBodyValid || err != nil {
		log.Printf("Erro ao validar o body: %v", err)
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		return
	}

	rows, err := models.Update(username, user)
	if err != nil {
		log.Printf("Erro ao atualizar o registro: %v", err)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	} else if rows == 0 {
		http.Error(w, "404 - Usuário não encontrado", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Usuário atualizado"))
	}

}
