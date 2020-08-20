package repository

import (
	"context"
	"go-echo-real-project/internal/model"
)

type UserRepository interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
}
