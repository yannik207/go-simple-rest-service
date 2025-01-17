package main

import (
	"task-api/server"
)

func main() {
	// Start HTTP server
	test := server.NewAPIServer(":8080")
	test.Run()
}
