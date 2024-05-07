package user

type User struct {
	Username string
	Password string
	ID       string
}

func (u *User) Login(username, password string) bool {
	if username == "testuser" && password == "testpassword" {
		return true
	}
	return false
}

func AddUser(username string) (User, error) {
	newUser := User{ID: "1", Username: username}
	return newUser, nil
}
