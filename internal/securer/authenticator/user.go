package authenticator

var users = map[string]bool{
	"catherine.kamkova@gmail.com": true,
	"arykalin@gmail.com":          true,
	"mishalesha@gmail.com":        true,
}

func UserIsValid(user string) bool {
	if valid, ok := users[user]; ok {
		return valid
	}
	return false
}
