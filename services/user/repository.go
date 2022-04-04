package user

import repository "trial-go-stacktrace/repository/user"

type userRepo interface {
	GetUser() error
}

func NewUserRepo() userRepo {
	userRepo := repository.UserImpl{}
	return &userRepo
}
