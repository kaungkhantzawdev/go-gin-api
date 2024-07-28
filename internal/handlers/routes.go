package handlers

import (
	"go-gin-api/internal/repository"
	"go-gin-api/internal/services"
	"github.com/gin-gonic/gin"
)

// UserRoutes registers the user-related routes.
func UserRoutes(router *gin.RouterGroup) {
	// Initialize UserRepository and UserService
	userRepo := &repository.UserRepository{}
	userService := services.NewUserService(userRepo)

	// Initialize the UserHandler
	userHandler := NewUserHandler(userService)

	// Register routes with userHandler
	router.GET("/users", userHandler.GetUsers)
}