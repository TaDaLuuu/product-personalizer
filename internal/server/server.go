package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go-echo-real-project/api"
	db2 "go-echo-real-project/internal/db"
	"go-echo-real-project/internal/handler"
	"go-echo-real-project/internal/repository/repo_impl"
	"go-echo-real-project/internal/utils"
)

type Server struct {
}

func (s *Server) Listen() {
	e := echo.New()

	db := db2.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "postgres",
		DbName:   "go_echo",
	}

	db.Connect()

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
	e.Logger.Fatal(e.Start(":1323"))
}
