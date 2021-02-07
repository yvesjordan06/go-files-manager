package application

import (
	"files_manager/configs"
	"github.com/kataras/iris/v12"
	"log"
)

var (
	/// Server is the application instance server
	Server = new(iris.Application)
)

func init() {
	log.Println("Server initiating")
	Server = iris.Default()
}

/// Start the server on port default 7777
func Start() {
	Server.Listen(configs.Application.HostAndPort())
}
