package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
) // press ctr+q here to see how testing works

// go test -cover
// we will get
//100% files as user.go does not contain any methods or function
//Try to add a mock function and you will get less % of files

func TestUserNotFound(t *testing.T) {

	user, err := GetUser(0)
	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err, "We were expecting an error when user id is 0 ")
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, "Not Found", err.Code)
	assert.EqualValues(t, err.Message, "User 0 was not found")

	/*
		if user != nil {
			t.Errorf("We were not expecting a user with id 0")
		}
		if err==nil {
			t.Errorf("We were expecting an error when user id is 0")
		}

		if err.StatusCode!= http.StatusNotFound {
			t.Errorf("We were expecting 404 when user is not found")
		}
	*/

	/****************************************************
	Hacking Coverage
	GetUser(0)


	************************************************** */

}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Diwakar", user.FirstName)
	assert.EqualValues(t, "Singh", user.LastName)
	assert.EqualValues(t, "diwakar@gmail.com", user.Email)

	/****************************************************
	Remove Everything above
	Hacking Coverage
	Write this only
	GetUser(123)

	After putting these 2 function call we will increase the coverage
	even without testing anything.
	************************************************** */

}
