package main

import (
	"files_manager/application"
	_ "files_manager/models" // Do not delete
	_ "files_manager/routes" // Do not delete
)

func main() {

	// Starts the main application
	application.Start()
}
