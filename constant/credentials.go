package constant

import "strings"

const Username = "test"
const Password = "test"
const UsernameKey = "username"
const PasswordKey = "password"

func Authenticate(username, password string) bool {
	if strings.EqualFold(username, Username) && strings.EqualFold(password, Password) {
		return true
	}
	return false
}
