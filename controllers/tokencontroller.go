package controllers

import (
	"net/http"

	"jwt-authentication/auth"
	"jwt-authentication/config"
	"jwt-authentication/models"

	"github.com/gin-gonic/gin"
)

/*	Here we define a simple struct that will essentially be what the endpoint would expect as the request body.
	This would contain the user’s email id and password.
*/
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(ctx *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	// check if email exists and password is correct
	record := config.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
}
