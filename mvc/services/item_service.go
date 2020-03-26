package services

import (
	"diwakarsingh052/micro/mvc/domain"
	"diwakarsingh052/micro/mvc/utils"
	"net/http"
)

type itemService struct {}

var ItemService itemService

func (i *itemService)GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement Me",
		StatusCode: http.StatusInternalServerError,
		Code:       "500",
	}

}
