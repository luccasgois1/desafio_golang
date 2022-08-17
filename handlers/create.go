package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/luccascleann/projeto_api/models"
)

func ValidateBody(w http.ResponseWriter, r *http.Request) (bool, error) {
	mapBodyValidPatterns := map[string]string{
		"id":         `"id"?\s*:?\s*\d+`,
		"username":   `"username"\s*?:\s*?"`,
		"firstname":  `"firstName"\s*?:\s*?"`,
		"lastname":   `"lastName"\s*?:\s*?"`,
		"email":      `"email"\s*?:\s*?"`,
		"password":   `"password"\s*?:\s*?"`,
		"phone":      `"phone"\s*?:\s*?"`,
		"userStatus": `"userStatus"?\s*:?\s*\d+`,
	}
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false, err
	}
	bodyString := string(body)
	// Reseta o r.Body para o buffer json
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	if bodyString == "" {
		http.Error(w, "400 - request não possui body", http.StatusBadRequest)
		return false, err
	}
	
	for userAttribute, pattern := range mapBodyValidPatterns {
		match, err := regexp.MatchString(pattern, bodyString)
		if err != nil || !match {
			http.Error(w, fmt.Sprintf("400 - O atributo %s não está no body", userAttribute), http.StatusBadRequest)
			return false, err
		}
	}
	return true, err
}

func Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if isBodyValid, err := ValidateBody(w, r); !isBodyValid || err != nil {
		log.Printf("Erro ao validar o body: %v", err)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
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
