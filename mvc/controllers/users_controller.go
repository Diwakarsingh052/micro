package controllers

import (
	"log"
	"net/http"
)

// Controller is in charge to get the user.
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	log.Printf("About to process User_id %v", userIdParam)

	// Command to check above stuff
	// curl localhost:8080/users?user_id=123

}
