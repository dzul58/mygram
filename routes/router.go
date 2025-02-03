package routes

import (
	"mygram/controllers"
	"mygram/middlewares"
	"mygram/repositories"
	"mygram/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepository)

	// Initialize controllers
	userController := controllers.NewUserController(userService)

	// User routes
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.PUT("/", middlewares.Authentication(), userController.Update)
		userRoutes.DELETE("/", middlewares.Authentication(), userController.Delete)
	}

	return router
}
