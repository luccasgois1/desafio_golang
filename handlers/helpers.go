package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func ValidateUserName(w http.ResponseWriter,username string) (bool, error) {
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
