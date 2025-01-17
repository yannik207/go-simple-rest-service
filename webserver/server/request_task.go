package server

import (
	"fmt"
	"net/http"
)

func (s *APIServer) TaskHandler(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case "GET":
		return s.Get(w, r)
	case "POST":
		return s.Post(w, r)
	// case "PUT":
	// 	fmt.Println("PUT")
	// case "DELETE":
	// 	fmt.Println("DELETE")
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}
