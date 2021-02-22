package controllers

import (
	"files_manager/middleware"
	"files_manager/models"
	"files_manager/utilities"
	"github.com/kataras/iris/v12"
	"net/http"
	"strings"
)

func LoginController(ctx iris.Context) {
	data := &struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	ctx.ReadJSON(data)

	data.Username = strings.TrimSpace(data.Username)
	data.Password = strings.TrimSpace(data.Password)

	if data.Username == "" || data.Password == "" {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": "Invalid entries", "suggestion": "Username or Password must not be empty"})
		return
	}

	user := new(models.User)

	r, err := user.Get(&models.User{Username: data.Username})
	println(r)
	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": "User not found", "suggestion": "Please check your email"})
		return
	}

	if !utilities.CompareHashPassword(data.Password, user.Password) {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": "Incorrect Credential", "suggestion": "Please check your password"})
		return
	}

	token := models.Token{
		UserID: user.ID,
		User:   *user,
	}

	_, err = token.Create()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": "Can't create login Token", "suggestion": "Please contact the system admin"})
		return
	}

	ctx.StopWithJSON(http.StatusCreated, token)
}

func RegisterController(ctx iris.Context) {
	var data struct {
		models.User
		Password string `json:"password,omitempty" validate:"required"`
	}

	err := ctx.ReadJSON(&data)

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": "Invalid entries", "suggestion": err.Error()})
		return
	}

	user := data.User
	user.Password = data.Password

	r, err := user.Create()
	println(r)
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": "Can not create user", "suggestion": err.Error()})
		return
	}

	ctx.StopWithJSON(http.StatusCreated, user)
}

func Logout(ctx iris.Context) {
	middleware.AuthRequired()(ctx)
	if ctx.IsStopped() {
		return
	}

	token := ctx.Values().Get("token").(*models.Token)

	token.Disabled = true

	_, err := token.Save()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "Please report the error to the admin"})
		return
	}

	ctx.StopWithStatus(http.StatusOK)

}
