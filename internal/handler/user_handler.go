package handler

import (
	"github.com/labstack/echo/v4"
	"go-echo-real-project/internal/model"
	"go-echo-real-project/internal/model/request"
	"net/http"
)

type UserHandler struct {
}

func (u *UserHandler) Welcome(c echo.Context) error {
	user := model.User{
		FullName: "Hai VCL",
		Email:    "haivcl@gmail.commmm",
	}
	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) SignUp(c echo.Context) error {
	// Bind data from request
	req := request.SignUpRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Validate data
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, req)
}

func (u *UserHandler) SignIn(c echo.Context) error {
	return nil
}
