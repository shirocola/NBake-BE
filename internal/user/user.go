package user

import (
	"encoding/json"
	"net/http"

	"github.com/shirocola/NBake-BE/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
	ID       string
	Role     string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HandleLogin handles user login requests
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var usr User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if usr.Login(usr.Username, usr.Password) {
		token, err := auth.GenerateToken(usr.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
		return
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
}

// Login checks the credentials and returns true if they are valid
func (u *User) Login(username, password string) bool {
	return username == "testuser" && password == "testpassword"
}

// AddUser creates a new user and returns it
func AddUser(username string) (User, error) {
	newUser := User{ID: "1", Username: username}
	return newUser, nil
}
