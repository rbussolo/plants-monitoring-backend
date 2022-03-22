package user

import (
	"errors"
	"net/mail"
)

func CreateNewUser(email string, password string) (User, error) {
	u := User{}

	// Email is required
	if email == "" {
		return u, errors.New("Email is required.")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return u, errors.New("Email invalid.")
	}

	// Password is required
	if password == "" {
		return u, errors.New("Password is required.")
	}

	if len(password) < 6 {
		return u, errors.New("Password must to have at least 6 characters.")
	}

	u = FindUserByEmail(email)
	if u.Id > 0 {
		return u, errors.New("Email already has been used.")
	}

	u = NewUser(email, password)

	return u, nil
}
