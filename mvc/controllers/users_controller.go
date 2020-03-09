package controllers

import (
	"diwakarsingh052/micro/mvc/services"
	"diwakarsingh052/micro/mvc/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

// Controller is in charge to get the user.
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User Id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)

		return
	}

	//We need to send this user_id to service in order to get data
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
