package server

import (
	"fmt"
	"net/http"
	"task-api/server/methods"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		methods.Get(w, r)
	case "POST":
		fmt.Println("POST")
	case "PUT":
		fmt.Println("PUT")
	case "DELETE":
		fmt.Println("DELETE")
	default:
		fmt.Println("User X tried to access API via fault method")
		methods.FaultMethod(w, r)
	}
}