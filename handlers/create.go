package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/luccascleann/projeto_api/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if isBodyValid, err := ValidateBody(w, r); !isBodyValid || err != nil {
		log.Printf("Erro ao validar o body: %v", err)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = models.Insert(user)
	if err != nil {
		log.Printf("Erro ao tentar inserir um usuario ao banco de dados: %v", err)
		http.Error(w, "500 - Erro ao tentar inserir um usuario ao banco de dados", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - usuario criado"))
	}
}
