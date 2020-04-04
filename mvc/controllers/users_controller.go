package controllers

import (
	"diwakarsingh052/micro/mvc/services"
	"diwakarsingh052/micro/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Controller is in charge to get the user.
func GetUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User Id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(http.StatusBadRequest,apiErr)
		return
	}

	//We need to send this user_id to service in order to get data
	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode,apiErr)
		return
	}
	c.JSON(http.StatusOK,user)


}
