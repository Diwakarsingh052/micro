package app

import (
	"diwakarsingh052/micro/mvc/controllers"
	"net/http"
)

// this app.go starts the web server

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
