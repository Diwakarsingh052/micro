package main

import "net/http"

func main() {

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello World"))
		if err != nil {
			panic(err)
		}
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}

}

//curl localhost:8081/hello
