package unitservice

import (
	"fmt"
	"log"

	"github.com/ankur-anand/mocking-demo/userdb"
)

// User encapsulate a user in the system.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

// Mock objects meet the interface requirements of,
// and stand in for, more complex real ones
type RegistrationPreChecker interface {
	userExists(string) bool
}

type regPreCheck struct{}

func (r regPreCheck) userExists(email string) bool {
	return userdb.UserExists(email)
}


func NewRegistrationPreChecker() RegistrationPreChecker {
	return regPreCheck{}
}

// RegisterUser will register a User if only User has not been previously
// registered.
func RegisterUser(user User, regPreCond RegistrationPreChecker) error {
	// check if user is already registered
	found := regPreCond.userExists(user.Email)
	if found {
		return fmt.Errorf("email '%s' already registered", user.Email)
	}
	// carry business logic and Register the user in the system
	log.Println(user)
	return nil
}
