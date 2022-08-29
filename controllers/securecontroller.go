package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*	The idea is that we would secure this endpoint so that only
	the requests having a valid JWT at the Request header will
	be able to access this.
*/
func Ping(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Pong"})
}
