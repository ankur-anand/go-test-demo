package unit_simple_service

import "testing"

func TestRegisterUser(t *testing.T) {
	user := User{
		Name:     "Ankur Anand",
		Email:    "anand@example.com",
		UserName: "anand",
	}

	err := RegisterUser(user)
	if err != nil {
		t.Fatal(err)
	}
}
