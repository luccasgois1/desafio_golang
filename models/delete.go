package models

import "github.com/luccascleann/projeto_api/db"

func Delete(username string) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM users WHERE username=$1`, username)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
