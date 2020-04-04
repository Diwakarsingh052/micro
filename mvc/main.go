// Entry Point of our application
package main

import "diwakarsingh052/micro/mvc/app"



// Press ctrl+q to see documentation of given stuff
func main() {

	app.StartApp() // calling function from app package. Stored Locally.

}

// curl localhost:8080/users/123
// we changed here how we run our app. check previous commit how it works.
