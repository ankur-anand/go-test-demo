package unit_simple_service

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
type userDB interface {
	userExist(string) bool
}

type uDB struct{}

func (u uDB) userExist(email string) bool {
	return userdb.UserExist(email)
}

var userDatabase userDB

func init() {
	userDatabase = uDB{}
}

// RegisterUser will register a User if only User has not been previously
// registered.
func RegisterUser(user User) error {
	// check if user is already registered
	found := userDatabase.userExist(user.Email)
	if found {
		return fmt.Errorf("email '%s' already registered", user.Email)
	}
	// carry business logic and Register the user in the system
	log.Println(user)
	return nil
}
