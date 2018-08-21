package controllers

import (
	"crypto/rsa"
	// "crypto/sha256"
	// jwt "github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"goauth/app"
	"goauth/models"
	"goauth/utils"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func SignIn(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		log.Fatalln("User bind json error", err)
		return
	}

	auth, err := models.FindUser(app.GetDB(c), &user)

	if err != nil {
		utils.FailedNotFound(c, gin.H{
			"status":  http.StatusNotFound,
			"message": "Login Failed",
			"data":    nil,
		})
	} else {
		auth.Token = utils.GenerateJWT(user)

		utils.SuccessOK(c, gin.H{
			"status":  http.StatusOK,
			"message": "Login Success",
			"data":    auth,
		})
	}
}

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		log.Fatalln("User bind json error", err)
		return
	}

	if err := user.Create(app.GetDB(c)); err != nil {
		log.Fatalln("Create User Model err", err)
		return
	}

	c.JSON(http.StatusCreated, user)
}
