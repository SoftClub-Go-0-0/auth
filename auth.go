package auth

var users = map[string]string{
	"paraparadox": "pass",
}

func Validate(username, password string) bool {
	for user, pass := range users {
		if username == user && password == pass {
			return true
		}
	}
	return false
}

func AddUser(username, password string) bool {
	if _, exists := users[username]; !exists {
		users[username] = password
		return true
	}
	return false
}

func RemoveUser(username string) bool {
	if _, exists := users[username]; exists {
		delete(users, username)
		return true
	}
	return false
}
