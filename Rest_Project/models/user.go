package models

import "example.com/rest-project/db"

type User struct {
	ID       int64
	Email    string `binding:required`
	Password string `binding:required`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password)
	VALUES ($1, $2)
	RETURNING id;
	`
	statement, statement_err := db.DB.Prepare(query)
	if statement_err != nil {
		return statement_err
	}
	defer statement.Close()

	var id_num int
	result_err := statement.QueryRow(u.Email, u.Password).Scan(&id_num)
	if result_err != nil {
		return result_err
	}

	u.ID = int64(id_num)
	return nil
}
