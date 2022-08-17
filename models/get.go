package models

import (
	"fmt"

	"github.com/luccascleann/projeto_api/db"
)

func Get(username string) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		fmt.Printf("Erro ao abrir o banco de dados: %v", err)
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM users WHERE username=$1`, username)
	err = row.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.Pass, &user.Phone, &user.Userstatus)

	return
}
