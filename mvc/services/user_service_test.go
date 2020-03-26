package services

import (
	"diwakarsingh052/micro/mvc/domain"
	"diwakarsingh052/micro/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var userDaoMock usersDaoMock
var getUserFunction func(userId int64) (*domain.User,*utils.ApplicationError)

type usersDaoMock struct{}
func (m *usersDaoMock) GetUser(userid int64) (*domain.User, *utils.ApplicationError){
	return getUserFunction(userid)
}

func TestGetUserNotFound(t *testing.T) {
	user,err := UserService.GetUser(0)
	assert.Nil(t,user)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t,"User 0 was not found",err.Message )
}
