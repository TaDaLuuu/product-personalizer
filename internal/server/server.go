package server

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-echo-real-project/api"
	"go-echo-real-project/internal/config"
	db2 "go-echo-real-project/internal/db"
	"go-echo-real-project/internal/handler"
	"go-echo-real-project/internal/repository/repo_impl"
	"go-echo-real-project/internal/utils"
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

	userRepository := repo_impl.NewUserRepository(&db)
	userHandler := handler.UserHandler{
		userRepository,
	}
	apiRouter := api.Api{
		Echo:        e,
		UserHandler: userHandler,
	}
	apiRouter.SetupRouter()
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.Port)))
}
