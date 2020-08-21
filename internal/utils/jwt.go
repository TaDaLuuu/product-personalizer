package utils

import (
	"github.com/dgrijalva/jwt-go"
	"go-echo-real-project/internal/model"
	"time"
)

// Read from .env
const SECRET_KEY = "VCL"

func GenToken(user model.User) (string, error) {
	claims := &model.JWTCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
