package main

import (
	"go-api/config"
	"go-api/controllers"
	"go-api/repos"
	"go-api/routes"
	"go-api/services"
)

func main() {
	
	config.ConnectDB()

	// Initialize repositories
	userRepo := repos.NewUserRepository()
	taskRepo := repos.NewTaskRepository()

	authService := services.NewAuthService(userRepo)
	taskService := services.NewTaskService(taskRepo)

	authController := controllers.NewAuthController(authService)
	taskController := controllers.NewTaskController(taskService)

	router := routes.SetupRoutes(authController, taskController)
	router.Run(":8080")
}
