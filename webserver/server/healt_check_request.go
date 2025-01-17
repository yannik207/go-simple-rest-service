package server

import (
	"net/http"
	"task-api/server/methods"
)

func (s *APIServer) WelcomeHandler(w http.ResponseWriter, r *http.Request) error {

	statusCode := http.StatusOK //200

	return methods.WriteJSON(w, statusCode, "Hello from the API!")
}
