package api

import (
	"github.com/labstack/echo/v4"
	"go-echo-real-project/internal/handler"
)

type Api struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *Api) SetupRouter() {
	api.Echo.GET("/", api.UserHandler.Welcome)

	authGroup := api.Echo.Group("/auth")
	authGroup.POST("/sign-in", api.UserHandler.SignIn)
	authGroup.POST("/sign-up", api.UserHandler.SignUp)
}
