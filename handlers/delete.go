package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luccascleann/projeto_api/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if isUserNameValid, err := ValidateUserName(w, username); !isUserNameValid || err != nil {
		log.Printf("Erro ao validar o username: %v", err)
		return
	}

	rows, err := models.Delete(username)
	if err != nil {
		log.Printf("Erro ao remover o registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	} else if rows == 0 {
		http.Error(w, "404 - Usuário não encontrado", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Usuário apagado"))
	}

}
