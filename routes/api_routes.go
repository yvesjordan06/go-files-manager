package routes

import "files_manager/application"

func init() {
	api := application.Server.Party("api/")
	authParty := api.Party("/auth")
	{
		authParty.Post("/login")
		authParty.Post("/register")
		authParty.Get("/user")
	}
}
