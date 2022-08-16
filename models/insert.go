package models

import (
	"fmt"

	"github.com/luccascleann/projeto_api/db"
)

func Insert(user User) (err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		fmt.Printf("Erro ao abrir o banco de dados: %v", err)
		return
	}
	defer conn.Close()
	
	sql_command := `INSERT INTO users (id, username, firstName, lastName, email, password, phone, userStatus) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = conn.Exec(sql_command, user.ID, user.Username, user.Firstname, user.Lastname, user.Email, user.Pass, user.Phone, user.Userstatus)

	return
}
