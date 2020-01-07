package userdb

import "testing"

func TestUserExist(t *testing.T) {
	found := UserExist("anand@example.com")
	if found != true {
		t.Errorf("expected to return true for existing user")
	}
}
