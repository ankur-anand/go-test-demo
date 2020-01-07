package simpleservice

import "testing"

func TestCheckUserExist(t *testing.T) {
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
