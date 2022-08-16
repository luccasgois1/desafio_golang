package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/luccascleann/projeto_api/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, "400 - request não possui body", http.StatusBadRequest)
		return
	}
	
	err = models.Insert(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao tentar inserir um usuario ao banco de dados: %v", err), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - usuario criado"))
	}
}
