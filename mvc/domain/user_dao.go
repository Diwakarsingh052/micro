package domain

import (
	"errors"
	"fmt"
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

func GetUser(userId int64) (*User, error) {
	/*
		user := person[userId]

		if user == nil { // if no user found with particular id
		return nil, errors.New(fmt.Sprintf("user %v not found", userId))
		}
		//We asked *User in argument and
		// that's the reason we can return nil above.
		// instead of any User type value

		return user, nil
	*/

	if user := person[userId]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("user %v not found", userId))

}
