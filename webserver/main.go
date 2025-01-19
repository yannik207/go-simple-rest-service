package main

import (
	"task-api/db"
	"task-api/server"
)

func main() {
	// Start HTTP server
	db.Conn()

	test := server.NewAPIServer(":8080")
	test.Run()
}
