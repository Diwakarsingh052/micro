package app

import (
	"diwakarsingh052/micro/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
	// by (colon):user_id we are defining user_id is a variable.
}
