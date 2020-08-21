package repository

import (
	"context"
	"go-echo-real-project/internal/model"
	"go-echo-real-project/internal/model/request"
)

type UserRepository interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	FindUser(context context.Context, request request.SignInRequest) (model.User, error)
}
