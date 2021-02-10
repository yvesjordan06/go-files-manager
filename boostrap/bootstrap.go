package boostrap

import (
	"files_manager/application"
	"files_manager/middleware"
	"log"
)

/// Bootstrap, Here we start service and task before running the server
/// It can be used to register global middlewares or plugin

func init() {
	log.Println("Starting bootstrap services")
	application.Server.Use(middleware.BasicMiddleware())
	application.Server.Use(middleware.CorsMiddleware())
}
