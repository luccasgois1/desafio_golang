package models

import "github.com/luccascleann/projeto_api/db"

func GetAll() (users []User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM users`)
	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.Pass, &user.Phone, &user.Userstatus)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	return
}
