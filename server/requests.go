package server

import (
	"fmt"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Type of ResponseWriter: %T\n", w)

	statusCode := http.StatusOK //200

	w.WriteHeader(statusCode)

	// Return a response to the user
	responseMessage := fmt.Sprintf("Status code: %d, Hello from the API!", statusCode)
	w.Write([]byte(responseMessage)) // Send the response message
}
