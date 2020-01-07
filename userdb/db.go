package userdb

// User encapsulate a user in the system.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

// db act as a dummy package level database.
var db map[string]User

// init initialize a dummy db with some data
func init() {
	db = make(map[string]User)
	db["ankuranand@dummy.org"] = User{
		Name:     "Ankur Anand",
		Email:    "ankuranand@dummy.org",
		UserName: "ankuranand",
	}
	db["anand@example.com"] = User{
		Name:     "Anand",
		Email:    "anand@example.com",
		UserName: "anand",
	}
}

// UserExist check if the User is registered with the provided email.
func UserExist(email string) bool {
	if _, ok := db[email]; !ok {
		return false
	}
	return true
}
