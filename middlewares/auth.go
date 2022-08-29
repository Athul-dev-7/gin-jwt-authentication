/*	So, how do we check if the incoming request contains a valid token?
	We need to place this validation check somewhere globally and make
	it usable by all the endpoints we need to secure.

	Middleware is the answer to this. So, what a Middleware does is,
	it attaches to the HTTP pipeline of the application.
	So, once a client sends a request, the first block it hits will
	be the middleware, only after this, the request will be hitting
	the actual endpoint.
	That’s a suitable place to position our Token Validation check.
*/

package middlewares

import (
	"net/http"

	"jwt-authentication/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extracts the Authorization header from the HTTP context.
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		// validate the token
		err := auth.ValidateToken(tokenString)
		if err != nil {
			// If the token is found to be invalid or expired, the application would throw a 401 Unauthorized exception.
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		// If the token is valid, the middleware allows the flow and the request reaches the required controller’s endpoint.
		ctx.Next()
	}
}
