package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/luccascleann/projeto_api/models"
)

func ValidateUserName(w http.ResponseWriter, username string) (bool, error) {
	invalidPatterns := `[!@#$%*(){}]`
	matched, err := regexp.MatchString(invalidPatterns, username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else if matched {
		http.Error(w, "400 - username inválido", http.StatusBadRequest)
	}
	return !matched, err
}

func ValidateBody(w http.ResponseWriter, r *http.Request) (bool, error) {
	mapBodyValidPatterns := map[string]string{
		"id":         `"id"?\s*:?\s*\d+`,
		"username":   `"username"\s*?:\s*?"`,
		"firstName":  `"firstName"\s*?:\s*?"`,
		"lastName":   `"lastName"\s*?:\s*?"`,
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
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return false, err
		}
		if !match {
			http.Error(w, fmt.Sprintf("400 - O atributo %s não está no body", userAttribute), http.StatusBadRequest)
			return false, err
		}
	}
	return true, err
}

func IsUserAlreadyRegisted(w http.ResponseWriter, users []models.User) bool {
	for _, user := range users {
		_, err := models.Get(user.Username)
		if err != sql.ErrNoRows {
			http.Error(w, fmt.Sprintf(`400 - O username "%s" ja esta cadastrado.`, user.Username), http.StatusBadRequest)
			return true
		}
	}
	return false
}

// TODO Join ValidateBody and ValidateBodyArray
func ValidateBodyArray(w http.ResponseWriter, r *http.Request) (bool, error) {
	mapBodyValidPatterns := map[string]string{
		"id":         `"id"?\s*:?\s*\d+`,
		"username":   `"username"\s*?:\s*?"`,
		"firstName":  `"firstName"\s*?:\s*?"`,
		"lastName":   `"lastName"\s*?:\s*?"`,
		"email":      `"email"\s*?:\s*?"`,
		"password":   `"password"\s*?:\s*?"`,
		"phone":      `"phone"\s*?:\s*?"`,
		"userStatus": `"userStatus"?\s*:?\s*\d+`,
	}

	// Armazena as informacoes do body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false, err
	}
	bodyString := string(body)
	// Reseta o r.Body para o buffer json
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// Testa se o body esta vazio
	if bodyString == ""{
		http.Error(w, "400 - request não possui body", http.StatusBadRequest)
		return false, err
	}

	// Valida se todos os usuarios nao existem no banco
	var users []models.User
	err = json.NewDecoder(r.Body).Decode(&users)
	// Reseta o r.Body para o buffer json
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return false, err
	}
	numberOfUsersOnBody := len(users)
	if numberOfUsersOnBody == 0{
		http.Error(w, "400 - request não possui body", http.StatusBadRequest)
		return false, err
	} 
	if IsUserAlreadyRegisted(w, users) {
		return false, err
	}

	// Verifica se esqueceram que colocar algum atributo no users do body
	for userAttribute, pattern := range mapBodyValidPatterns {
		re := regexp.MustCompile(pattern)
		numberOfMatches := len(re.FindAllString(bodyString, -1))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return false, err
		}
		if numberOfMatches < numberOfUsersOnBody {
			http.Error(w, fmt.Sprintf("400 - O atributo %s não está no body", userAttribute), http.StatusBadRequest)
			return false, err
		}
	}
	return true, err
}
