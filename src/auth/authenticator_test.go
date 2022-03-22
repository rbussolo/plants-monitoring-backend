package auth

import (
	"testing"

	u "backend/src/user"
)

func TestAuth(t *testing.T) {
	user, err := getApiKeyTestUser()

	if err != nil {
		t.Fatalf("error trying create a new user test %s", err.Error())
	}

	// Try to authenticate
	token, err := Auth(user.ApiKey, user.Email)

	if err != nil {
		t.Fatalf("error trying create a new token %s", err.Error())
	}

	// Try to decode token
	isAuth, user_id := IsAuthenticated(token)

	if !isAuth {
		t.Fatal("token isn't authenticated")
	}

	if user_id != user.Id {
		t.Fatalf("excepted %d user id, but receive %d", user.Id, user_id)
	}
}

func getApiKeyTestUser() (u.User, error) {
	var email string = "test_authenticator@test.com"
	var password string = "test_auth"
	var err error

	// Try to find test user
	user := u.FindUserByEmail(email)

	if user.Id == 0 { // Create user test
		user, err = u.CreateNewUser(email, password)
	}

	return user, err
}
