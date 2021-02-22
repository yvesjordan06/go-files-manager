package middleware

import (
	"files_manager/models"
	"github.com/kataras/iris/v12"
	"log"
	"net/http"
	"strings"
)

type tokenHeader struct {
	Token string `header:"Authorization,required"`
}

func AuthRequired() iris.Handler {
	return func(ctx iris.Context) {
		var authHeader tokenHeader

		// Trying to get headers
		if err := ctx.ReadHeaders(&authHeader); err != nil {
			ctx.StopWithJSON(http.StatusUnauthorized, iris.Map{"error": err.Error(), "suggestion": "We could not find Authorization as part of your headers, Please add it with your token"})
			return
		}

		// Checking if the token is in the database and valid
		token := strings.TrimSpace(strings.Split(authHeader.Token, "Bearer")[1])
		log.Println("User token is :", token)
		dbtoken := new(models.Token)
		_, err := dbtoken.Get(&models.Token{Token: token})

		if err != nil || dbtoken.Disabled {
			ctx.StopWithJSON(http.StatusUnauthorized, iris.Map{"error": "Invalid Token", "suggestion": "Your token could not be found or may have been invalidated, please login to continue"})
			return
		}

		log.Println(dbtoken)

		ctx.SetUser(dbtoken.User)
		ctx.Values().Set("token", dbtoken)
		ctx.Values().Set("user", &(dbtoken.User))

		ctx.Next()
	}
}
