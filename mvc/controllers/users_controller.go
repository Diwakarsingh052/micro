package controllers

import (
	"diwakarsingh052/micro/mvc/services"
	"encoding/json"
	"net/http"
	"strconv"
)

// Controller is in charge to get the user.
func GetUser(resp http.ResponseWriter, req *http.Request) {
	/*
		Long Way
		userIdParam := req.URL.Query().Get("user_id")
		userId, err := strconv.ParseInt(userIdParam),10,64)
	*/
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {

		//Added after creating jsonValue

		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("User Id must be a number"))
		// till here
		return
	}

	//We need to send this user_id to service in order to get data
	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
