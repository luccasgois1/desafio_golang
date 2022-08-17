package models

import (
	"fmt"

	"github.com/luccascleann/projeto_api/db"
)

func Update(username string, user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		fmt.Printf("Erro ao abrir o banco de dados: %v", err)
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE users SET id=$1, username=$2, firstName=$3, lastName=$4, email=$5, password=$6, phone=$7, userStatus=$8 WHERE username=$9`,
		user.ID, user.Username, user.Firstname, user.Lastname, user.Email, user.Pass, user.Phone, user.Userstatus, username)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
