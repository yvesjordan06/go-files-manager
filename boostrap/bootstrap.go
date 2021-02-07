package boostrap

import (
	"files_manager/application"
	"files_manager/middleware"
)

/// Bootstrap, Here we start service and task before running the server
/// It can be used to register global middlewares or plugin

func init() {
	application.Server.Use(middleware.BasicMiddleware())
}
