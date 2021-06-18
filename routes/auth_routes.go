package routes

import (
	"files_manager/controllers"
	"files_manager/middleware"

	"github.com/kataras/iris/v12"
)

func init() {
	authParty := api.Party("/auth")
	{
		authParty.Post("/login", controllers.LoginController)
		authParty.Post("/logout", controllers.Logout)
		authParty.Post("/register", controllers.RegisterController)
		authParty.Get("/user", middleware.AuthRequired(), func(ctx iris.Context) {
			ctx.StopWithJSON(200, ctx.Values().Get("user"))
		})
	}
}
