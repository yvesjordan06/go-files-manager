package routes

import "files_manager/application"

func init() {
	authParty := application.Server.Party("/auth")
	{
		authParty.Post("/login")
		authParty.Post("/register")
		authParty.Get("/user")
	}
}
