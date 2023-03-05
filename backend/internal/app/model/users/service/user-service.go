package service

import (
	"github.com/sirupsen/logrus"
	"prj/internal/app/model/users/entity"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
}

type userService struct {
	users  []entity.User
	logger *logrus.Logger
}

func New() UserService {
	return &userService{
		logger: logrus.New(),
	}
}

func (service *userService) Save(user entity.User) entity.User {
	service.logger.Info("Save User: ", user)
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindAll() []entity.User {
	service.logger.Info("Find All Users")
	return service.users
}
