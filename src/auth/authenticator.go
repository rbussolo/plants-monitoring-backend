package auth

import (
	"errors"
	"net/mail"
	"time"

	u "backend/src/user"

	"github.com/dgrijalva/jwt-go"
)

const access_secret = "2b93cda3-bdb8-4c1c-8467-fc7c09637f23"

func Auth(apiKey string, email string) (string, error) {
	var token string

	if apiKey == "" {
		return token, errors.New("Apikey is required.")
	}

	if email == "" {
		return token, errors.New("Email is required.")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return token, errors.New("Email invalid.")
	}

	user := u.FindUserByEmail(email)
	if user.Id == 0 {
		return token, errors.New("User not found.")
	}

	if user.ApiKey != apiKey {
		return token, errors.New("Apikey invalid.")
	}

	token, err = createToken(user.Id)
	if err != nil {
		return token, err
	}

	return token, nil
}

func AuthWithPassword(email string, password string) (string, error) {
	var token string

	if email == "" {
		return token, errors.New("Email is required.")
	}

	if password == "" {
		return token, errors.New("Password is required.")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return token, errors.New("Email invalid.")
	}

	user := u.FindUserByEmailPassword(email, password)
	if user.Id == 0 {
		return token, errors.New("Email / password invalid.")
	}

	// Create token to access
	token, err = createToken(user.Id)

	return token, err
}

func IsAuthenticated(token string) (bool, int) {
	var user_id int
	var exp int64

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(access_secret), nil
	})

	if err != nil {
		return false, user_id
	}

	for key, val := range claims {
		if key == "user_id" {
			user_id = int(val.(float64))
		} else if key == "exp" {
			exp = int64(val.(float64))
		}
	}

	tm := time.Unix(exp, 0)
	if time.Now().After(tm) {
		return false, user_id
	}

	return true, user_id
}

func createToken(user_id int) (string, error) {
	var err error

	// Define payload
	claims := jwt.MapClaims{}
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Define signing method
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign string with secret key
	token, err := at.SignedString([]byte(access_secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
