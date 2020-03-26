package services

import (
	"diwakarsingh052/micro/mvc/domain"
	"diwakarsingh052/micro/mvc/utils"
)

type userService struct {

}
var UserService userService

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {

	return domain.UserDao.GetUser(userId)



}