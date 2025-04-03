package routes

import (
	"net/http"
	"time"

	"github.com/carlosgrillet/go-api/models"
	"github.com/carlosgrillet/go-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format"})
		return
	}
 
  _, err = models.GetUserByEmail(user.Email)
  if err != nil {
    user.ID = utils.GenerateID()
    user.CreatedAt = time.Now()
    go user.Save()
    context.JSON(http.StatusOK, gin.H{"message": "User created"})
		return
  }

  context.JSON(http.StatusInternalServerError, gin.H{"error": "User registred"})
}

func login(context *gin.Context) {
  var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request format"})
		return
	}

  userId, err := user.ValidateCredentials() 
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

  token, err := utils.GenerateToken(userId, user.Email)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  
  context.JSON(http.StatusOK, gin.H{"access_token": token})
}
