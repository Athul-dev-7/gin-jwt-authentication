package controllers

import (
	"net/http"

	"jwt-authentication/config"
	"jwt-authentication/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	record := config.DB.Create(&user)
	if record.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		ctx.Abort()
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
