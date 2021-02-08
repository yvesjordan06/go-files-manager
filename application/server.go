package application

import (
	"files_manager/configs"
	"github.com/kataras/iris/v12"
	"log"
	"strings"
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
	if configs.Application.AutoTLS {
		Server.Run(iris.AutoTLS(configs.Application.SecureAddress(), strings.Join(configs.Application.Domains, " "), strings.Join(configs.Application.Emails, " ")))
	}
	Server.Listen(configs.Application.DefaultAddress())
}
