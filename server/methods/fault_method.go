package methods

import (
	"fmt"
	"net/http"
)


func FaultMethod(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusMethodNotAllowed //405

	w.WriteHeader(statusCode)

	// Return a response to the user
	responseMessage := fmt.Sprintf("Method is not Allowed! %d", statusCode)
	w.Write([]byte(responseMessage)) // Send the response message
}