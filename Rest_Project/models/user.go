package models

import (
	"errors"

	"example.com/rest-project/db"
	"example.com/rest-project/utils"
)

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
	//Hash the password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	var id_num int
	result_err := statement.QueryRow(u.Email, hashedPassword).Scan(&id_num)
	if result_err != nil {
		return result_err
	}

	u.ID = int64(id_num)
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := db.DB.QueryRow(query, u.Email)

	var retrievePassword string
	err := row.Scan(&u.ID, &retrievePassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievePassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
