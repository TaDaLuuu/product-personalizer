package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-echo-real-project/internal/model"
	"go-echo-real-project/internal/model/request"
	"go-echo-real-project/internal/repository"
	"go-echo-real-project/internal/utils"
	"net/http"
	"time"
)

type UserHandler struct {
	UserRepository repository.UserRepository
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

	// Save user
	hashedPassword := utils.HashAndSalt([]byte(req.Password))
	userId := uuid.New()
	user := model.User{
		UserId:    userId.String(),
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      model.MEMBER.String(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Token:     "",
	}
	user, err := u.UserRepository.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng ký thành công",
		Data:       user,
	})
}

func (u *UserHandler) SignIn(c echo.Context) error {
	// Bind data from request
	req := request.SignInRequest{}
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

	user, err := u.UserRepository.FindUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	isPasswordValidated := utils.ComparePasswords(user.Password, []byte(req.Password))
	if !isPasswordValidated {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Sai mật khẩu",
			Data:       nil,
		})
	}

	token, err := utils.GenToken(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng nhập thành kông",
		Data:       user,
	})
}

func (u *UserHandler) Profile(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JWTCustomClaims)

	user, err := u.UserRepository.FindUserByEmail(c.Request().Context(), claims.Email)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "",
		Data:       user,
	})
}
