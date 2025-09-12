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
	query := `
		INSERT INTO users(email, password) 
		VALUES ($1, $2) 
		RETURNING id`

	hashedPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	err = database.DB.QueryRow(query, u.Email, hashedPass).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password from users WHERE email = $1"

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
