package user

import (
	"crypto/sha256"
	"encoding/base64"

	"database/sql"

	"github.com/google/uuid"

	d "backend/src/database"
)

type User struct {
	Id     int    `json:"id"`
	Email  string `json:"email"`
	ApiKey string `json:"apikey"`
}

func NewUser(email string, password string) User {
	db := d.GetInstance()

	var id int
	apiKey := uuid.New().String()

	// Password must to be encrypted
	encryottedPassword := encryptPassword(password)

	// Create a new user
	err := db.QueryRow("INSERT INTO users(email, password, apikey) VALUES($1, $2, $3) RETURNING id", email, encryottedPassword, apiKey).Scan(&id)
	if err != nil {
		panic(err)
	}

	u := User{
		Id:     id,
		Email:  email,
		ApiKey: apiKey,
	}

	return u
}

func FindUserByEmail(email string) User {
	db := d.GetInstance()

	var user User
	err := db.Get(&user, "SELECT id, email, apikey FROM users WHERE email = $1", email)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return user
}

func FindUserByEmailPassword(email string, password string) User {
	db := d.GetInstance()

	encryottedPassword := encryptPassword(password)

	var user User
	err := db.Get(&user, "SELECT id, email, apikey FROM users WHERE email = $1 AND password = $2", email, encryottedPassword)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return user
}

func encryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}
