package api

import (
	"github.com/labstack/echo/v4"
)

type Api struct {
	Echo *echo.Echo
}

func (api *Api) SetupRouter() {
}
