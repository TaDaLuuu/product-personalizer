package model

import "github.com/dgrijalva/jwt-go"

type JWTCustomClaims struct {
	jwt.StandardClaims
	UserId string
	Email  string
	Role   string
}
