package services

import (
	"diwakarsingh052/micro/mvc/domain"
	"diwakarsingh052/micro/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {

	// service will get data from domain

	return domain.GetUser(userId)

	// above statement will get replaced by return values of user_services.go

	//After the response we send it back to controller and it is
	//incharge to display the info see the mvc chart (refer)

}