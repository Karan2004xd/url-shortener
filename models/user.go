package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"url-shortner/db"
	"url-shortner/internal"
)

type User struct {
	Id int64 `json:"id"`
	Email string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (user *User) Copy(other User) {
	user.Id = other.Id
	user.Email = other.Email
	user.Password = other.Password
}

func (user *User) Create() error {
	hashedPassword, err := internal.HashPassword(user.Password)

	if err != nil {
		return errors.New(fmt.Sprint("Error while encrypting the password.", err))
	}

	result, err := db.Insert(db.InsertNewUser, user.Email, hashedPassword)

	if err != nil {
		return errors.New(fmt.Sprint("Error while saving the user.", err))
	}

	user.Id, err = result.LastInsertId()

	if err != nil {
		return errors.New(fmt.Sprint("Error while fetching the last inserted row.", err))
	}
	return nil
}

func createUserArrayFromRows(rows *sql.Rows, showPassword bool) ([]User, error) {
	var users []User

	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id, &user.Email,
			&user.Password, &user.CreatedAt)

		if err != nil {
			return nil, err
		}

		if !showPassword {
			user.Password = ""
		}

		users = append(users, user)
	}

	return users, nil
}

func (user *User) Validate() (error) {
	rows, err := db.Select(
		db.SelectUserByEmailAndPassword,
		user.Email)

	if err != nil {
		return err
	}

	users, err := createUserArrayFromRows(rows, true)

	if err != nil {
		return err
	}

	if len(users) == 0 {
		return errors.New("Invalid Credentials")
	}

	if !internal.CheckHashedPassword(users[0].Password, user.Password) {
		return errors.New("Invalid Credentials")
	}

	user.Id = users[0].Id
	return nil
}

func GetAllUsers() ([]User, error) {
	rows, err := db.Select(db.SelectAllUsers)

	if err != nil {
		return []User {}, err
	}

	users, err := createUserArrayFromRows(rows, false)

	if err != nil || len(users) == 0 {
		return []User {}, err
	}
	return users, nil 
}
