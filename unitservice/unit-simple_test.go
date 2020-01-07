package unitservice

import "testing"

// This helps in assigning mock at the runtime instead of compile time
var userExistMock func(email string) bool

type preCheckMock struct{}

func (u preCheckMock) userExists(email string) bool {
	return userExistMock(email)
}

func TestRegisterUser(t *testing.T) {
	user := User{
		Name:     "Ankur Anand",
		Email:    "anand@example.com",
		UserName: "anand",
	}

	regPreCond = preCheckMock{}
	userExistMock = func(email string) bool {
		return false
	}

	err := RegisterUser(user)
	if err != nil {
		t.Fatal(err)
	}

	userExistMock = func(email string) bool {
		return true
	}
	err = RegisterUser(user)
	if err == nil {
		t.Error("Expected Register User to throw and error got nil")
	}
}
