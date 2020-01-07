package unit_simple_service

import "testing"

// This helps in assigning mock at the runtime instead of compile time
var userExistMock func(email string) bool

type preCheckMock struct{}

func (u preCheckMock) userExist(email string) bool {
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
}
