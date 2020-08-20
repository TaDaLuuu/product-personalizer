package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-echo-real-project/api"
	"go-echo-real-project/internal/handler"
	"go-echo-real-project/internal/utils"
)

type Server struct {
}

func (s *Server) Listen() {
	e := echo.New()
	userHandler := handler.UserHandler{}
	apiRouter := api.Api{
		Echo:        e,
		UserHandler: userHandler,
	}
	apiRouter.SetupRouter()
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.Fatal(e.Start(":1323"))
}
