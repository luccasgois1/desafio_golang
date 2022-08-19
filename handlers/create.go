package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lib/pq"
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
	fmt.Println()
	if err != nil {
		pgErrorMessage := err.(*pq.Error).Message
		isPrimaryKeyInDataBase := strings.Contains(pgErrorMessage, "duplicate key value violates unique constraint")
		if isPrimaryKeyInDataBase {
			http.Error(w, fmt.Sprintf("400 - O usuario %s ja existe no banco de dados.", user.Username), http.StatusInternalServerError)
		} else {
			log.Printf("Erro ao tentar inserir um usuario ao banco de dados: %v", err)
			http.Error(w, "500 - Erro ao tentar inserir um usuario ao banco de dados", http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - usuario criado"))
	}
}
