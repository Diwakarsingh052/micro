package domain

import (
	"diwakarsingh052/micro/mvc/utils"
	"fmt"
	"net/http"
)

var (
	person = map[int64]*User{
		123: &User{
			Id:        1,
			FirstName: "Diwakar",
			LastName:  "Singh",
			Email:     "diwakar@gmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {


	if user := person[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found",userId),
		StatusCode: http.StatusNotFound,
		Code:       "Not Found",
	}

}
