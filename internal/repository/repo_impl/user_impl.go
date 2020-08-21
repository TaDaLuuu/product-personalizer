package repo_impl

import (
	"context"
	"go-echo-real-project/internal/db"
	"go-echo-real-project/internal/model"
	"go-echo-real-project/internal/model/request"
	"go-echo-real-project/internal/repository"
	"time"
)

type UserRepositoryImpl struct {
	sql *db.Sql
}

func (u UserRepositoryImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, full_name, email, password, role, created_at, updated_at)
		VALUES(:user_id, :full_name, :email, :password, :role, :created_at, :updated_at)
	`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u UserRepositoryImpl) FindUser(context context.Context, signInRequest request.SignInRequest) (model.User, error) {
	user := model.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE email = $1", signInRequest.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewUserRepository(sql *db.Sql) repository.UserRepository {
	return UserRepositoryImpl{sql}
}
