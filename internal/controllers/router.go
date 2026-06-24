package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/DanilAiro/kitty-app-back/internal/models"
	"github.com/DanilAiro/kitty-app-back/internal/repository"
	"github.com/DanilAiro/kitty-app-back/internal/service"
	"github.com/DanilAiro/kitty-app-back/internal/utils"
)

const (
	NoUserData          string = "NUD"
	BadUserEmail        string = "BUE"
	BadUserPassword     string = "BUP"
	CanNotCreateToken   string = "CNCT"
	CanNotCreateUser    string = "CNCU"
	BadToken            string = "BT"
	CanNotVerifyToken   string = "CNVT"
	UserExists          string = "UE"
	UserDoesNotExists   string = "UDNE"
	CanNotGenerateKitty string = "CNGK"
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
	var newUser *models.User
	if err := c.BindJSON(&newUser); err != nil {
		fail(c, http.StatusUnprocessableEntity, NoUserData, "no user data")
		return
	}

	if newUser.Email == "" {
		fail(c, http.StatusUnprocessableEntity, BadUserEmail, "no email")
		return
	}

	if newUser.Password == "" {
		fail(c, http.StatusUnprocessableEntity, BadUserPassword, "no password")
		return
	}

	savedUser := repository.GetUser(newUser)
	if savedUser != nil {
		fail(c, http.StatusConflict, UserExists, "user exists")
		return
	}

	password := []byte(newUser.Password)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fail(c, http.StatusInternalServerError, CanNotCreateUser, "can not create user")
		return
	}
	newUser.Password = string(hash)

	err = repository.AddUser(newUser)
	if err != nil {
		fail(c, http.StatusInternalServerError, CanNotCreateUser, "can not create user")
		return
	}

	ok(c, http.StatusCreated, gin.H{"message": "account created"})
}

func login(c *gin.Context) {
	var newUser *models.User
	if err := c.BindJSON(&newUser); err != nil {
		fail(c, http.StatusUnprocessableEntity, NoUserData, "no user data")
		return
	}

	if newUser.Email == "" {
		fail(c, http.StatusUnprocessableEntity, BadUserEmail, "no email")
		return
	}

	if newUser.Password == "" {
		fail(c, http.StatusUnprocessableEntity, BadUserPassword, "no password")
		return
	}

	savedUser := repository.GetUser(newUser)
	if savedUser == nil {
		fail(c, http.StatusConflict, UserDoesNotExists, "user does not exists")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(newUser.Password))
	if err != nil {
		fail(c, http.StatusUnauthorized, BadUserPassword, "password does not match")
		return
	}

	token, err := utils.CreateToken(savedUser.Email)
	if err != nil {
		fail(c, http.StatusInternalServerError, CanNotCreateToken, "can not create token")
		return
	}

	ok(c, http.StatusOK, gin.H{"jwt": token})
}

func kitty(c *gin.Context) {
	authString := c.Request.Header.Get("Authorization")
	rawToken := strings.Split(authString, " ")

	if len(rawToken) != 2 || strings.ToLower(rawToken[0]) != "bearer" {
		fail(c, http.StatusUnprocessableEntity, BadToken, "bad token")
		return
	}

	token := rawToken[1]

	_, err := utils.VerifyToken(token)
	if err != nil {
		fail(c, http.StatusUnauthorized, CanNotVerifyToken, "can not verify token")
		return
	}

	// добавить выдачу картинки
	kittyPhoto, err := service.GetKittyPhoto()
	if err != nil {
		fail(c, http.StatusInternalServerError, CanNotGenerateKitty, "can not generate kitty")
		return
	}

	ok(c, http.StatusOK, gin.H{"cat_url": kittyPhoto})
}
