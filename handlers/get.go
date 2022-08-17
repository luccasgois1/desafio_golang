package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
	"github.com/luccascleann/projeto_api/models"
)

func ValidateUserName(username string) (bool, error) {
	invalidPatterns := `[!@#$%*(){}]`
	matched, err := regexp.MatchString(invalidPatterns, username)
	return !matched, err
}

func Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	isUserNameValid, err := ValidateUserName(username)
	if err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}else if !isUserNameValid {
		http.Error(w, "400 - username inválido", http.StatusBadRequest)
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
