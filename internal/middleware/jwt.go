package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-echo-real-project/internal/model"
	"go-echo-real-project/internal/utils"
)

func CustomJWTMiddleware() echo.MiddlewareFunc {
	jwtConfig := middleware.JWTConfig{
		Claims:     &model.JWTCustomClaims{},
		SigningKey: []byte(utils.SECRET_KEY),
	}
	return middleware.JWTWithConfig(jwtConfig)
}
