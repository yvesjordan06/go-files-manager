package routes

import (
	"files_manager/controllers"
)

func init() {
	authParty := api.Party("/auth")
	{
		authParty.Post("/login", controllers.LoginController)
		authParty.Post("/logout", controllers.Logout)
		authParty.Post("/register", controllers.RegisterController)
		authParty.Get("/user")
	}
}
