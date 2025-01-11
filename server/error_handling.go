package server

import (
	"fmt"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusNotFound //404

	w.WriteHeader(statusCode)

	// Return a response to the user
	responseMessage := fmt.Sprintf("Status code: %d, Path not Found!", statusCode)
	w.Write([]byte(responseMessage)) // Send the response message
}