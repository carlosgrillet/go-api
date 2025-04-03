package routes

import (
	"github.com/carlosgrillet/go-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
  // Unprotected routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
  server.POST("/signup", signup)
  server.POST("/login", login)

  // Protected routes
  authenticated := server.Group("/")
  authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", addEvent)
  authenticated.PUT("/events/:id", updateEvent)
  authenticated.DELETE("/events/:id", deleteEvent)
}
