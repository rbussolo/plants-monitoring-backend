package user

import "testing"

type TestCaseNewUser struct {
	Email       string
	Password    string
	Fail        bool
	FailMessage string
}

func TestCreateNewUser(t *testing.T) {
	testCases := []TestCaseNewUser{
		{
			Email:    "rbussolo91@gmail.com",
			Password: "123456",
			Fail:     false,
		},
		{
			Email:       "outro_email",
			Password:    "123456",
			Fail:        true,
			FailMessage: "email invalid",
		},
		{
			Email:       "rbussolo91@gmail.com",
			Password:    "123",
			Fail:        true,
			FailMessage: "password must to have at least 6 characters",
		},
		{
			Email:       "rbussolo91@gmail.com",
			Password:    "123123",
			Fail:        true,
			FailMessage: "email already has been used",
		},
		{
			Email:    "carlos@gmail.com",
			Password: "123123",
			Fail:     false,
		},
	}

	for _, tc := range testCases {
		user, err := CreateNewUser(tc.Email, tc.Password)

		if err != nil && !tc.Fail {
			t.Fatalf("has expected success but receive error %s at user %s", err.Error(), user.Email)
		}

		if err == nil && tc.Fail {
			t.Fatalf("has expected fail but receive success at user %s", user.Email)
		}

		if err != nil && tc.Fail && tc.FailMessage != err.Error() {
			t.Fatalf("has expected %s but receive %s at user %s", tc.FailMessage, err.Error(), user.Email)
		}
	}
}
