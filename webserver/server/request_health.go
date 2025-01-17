package server

import (
	"net/http"
)

func (s *APIServer) WelcomeHandler(w http.ResponseWriter, r *http.Request) error {

	statusCode := http.StatusOK //200

	return WriteJSON(w, statusCode, "Hello from the API!")
}
