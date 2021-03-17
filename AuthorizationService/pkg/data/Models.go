package data

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	Id       int64  `json:"id" sql:"user_id"`
	Username string `json:"username" sql:"username"`
	Password string `json:"password" sql:"user_password"`
	Email    string `json:"email" sql:"email"`
}

type Login struct {
	Id       int64  `json:"id" sql:"user_id"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Users []*User

func HashPassword(user *Register) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
