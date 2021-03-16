package data

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserID int64
	Name   string
	Email  string
	*jwt.StandardClaims
}
