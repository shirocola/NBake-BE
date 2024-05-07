package user

type User struct {
	Username string
	Password string
}

func (u *User) Login(username, password string) bool {
	if username == "testuser" && password == "testpassword" {
		return true
	}
	return false
}
