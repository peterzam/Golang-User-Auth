package models

import jwt "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID uint
	Email  string
	Name   string
	Role   string
	Verify string
	*jwt.StandardClaims
}
