package routes

import (
	"go-api/controllers"
	"go-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(authController *controllers.AuthController, taskController *controllers.TaskController) *gin.Engine {
	router := gin.Default()

	// Auth Routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	// Task Routes (secured with JWT middleware)
	task := router.Group("/api/tasks")
	task.Use(middlewares.AuthMiddleware())
	{
		task.POST("/", taskController.CreateTask)
		task.GET("/", taskController.GetTasks)
		task.GET("/:id", taskController.GetTaskByID)
		task.PATCH("/:id", taskController.UpdateTask)
		task.DELETE("/:id", taskController.DeleteTask)
	}

	return router
}
