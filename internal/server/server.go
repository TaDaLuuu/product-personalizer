package server

import (
	"fmt"
	"go-echo-real-project/api"
	"go-echo-real-project/internal/config"
	db2 "go-echo-real-project/internal/db"
	"go-echo-real-project/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config *config.Config
}

func (s *Server) AutoInject() {
	config := config.LoadFromEnv()
	s.config = config
}

func (s *Server) Listen() {
	e := echo.New()

	db := db2.Sql{}
	db.Connect(s.config.Postgres)

	apiRouter := api.Api{
		Echo: e,
	}
	apiRouter.SetupRouter()
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.Port)))
}
