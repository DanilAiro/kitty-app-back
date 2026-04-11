package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	models "github.com/DanilAiro/kitty-app-back/internal/models"
	utils "github.com/DanilAiro/kitty-app-back/internal/utils"
)

const (
	NoUserData        string = "NUD"
	BadUserEmail      string = "BUE"
	BadUserPassword   string = "BUP"
	CanNotCreateToken string = "CNCT"
	BadToken          string = "BT"
	CanNotVerifyToken string = "CNVT"
)

func Router() {
	user := gin.Default()

	user.POST("/auth/register/", register)
	user.POST("/auth/login/", login)
	user.GET("/kitty/", kitty)

	user.Run()
}

func ok(c *gin.Context, status int, data interface{}) {
	c.JSON(status, models.Response{
		Success: true,
		Data:    data,
	})
}

func fail(c *gin.Context, status int, code, message string) {
	c.JSON(status, models.Response{
		Success: false,
		Error:   &models.ErrorInfo{Code: code, Message: message},
	})
}

func register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		fail(c, http.StatusNotFound, NoUserData, err.Error())
		return
	}

	if user.Email == "" {
		fail(c, http.StatusNotFound, BadUserEmail, "no email")
		return
	}

	if user.Password == "" {
		fail(c, http.StatusNotFound, BadUserPassword, "no password")
		return
	}

	// проверить наличие в БД

	// захешировать пароль

	// добавить запись в БД

	ok(c, http.StatusCreated, gin.H{"message": "account created"})
}

func login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		fail(c, http.StatusNotFound, NoUserData, err.Error())
		return
	}

	if user.Email == "" {
		fail(c, http.StatusNotFound, BadUserEmail, "no email")
		return
	}

	if user.Password == "" {
		fail(c, http.StatusNotFound, BadUserPassword, "no password")
		return
	}

	// проверить наличие в БД

	// захешировать пароль

	// сравнить пароли

	token, err := utils.CreateToken(user.Email)
	if err != nil {
		fail(c, http.StatusNotFound, CanNotCreateToken, err.Error())
		return
	}

	ok(c, http.StatusOK, gin.H{"jwt": token})
}

func kitty(c *gin.Context) {
	authString := c.Request.Header.Get("Authorization")
	rawToken := strings.Split(authString, " ")

	if len(rawToken) != 2 || strings.ToLower(rawToken[0]) != "bearer" {
		fail(c, http.StatusNotFound, BadToken, "bad token")
		return
	}

	token := rawToken[1]

	_, err := utils.VerifyToken(token)
	if err != nil {
		fail(c, http.StatusNotFound, CanNotVerifyToken, err.Error())
		return
	}

	// добавить выдачу картинки

	ok(c, http.StatusOK, gin.H{"jwt": token})
}
