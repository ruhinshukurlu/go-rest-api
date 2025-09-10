package models

import (
	"errors"

	"github.com/go-rest-api/database"
	"github.com/go-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string
	Password string
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password from users WHERE email = ?"

	row := database.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
