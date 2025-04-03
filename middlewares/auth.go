package middlewares

import (
	"net/http"

	"github.com/carlosgrillet/go-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
  token := context.Request.Header.Get("Authorization")

  if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token required"})
    return
  }

  userId, err := utils.ValidateToken(token)

  if err != nil {
    context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
    return
  }

  context.Set("userId", userId)
  context.Next()
}
