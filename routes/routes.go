package routes

import (
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
    "backend-onboard-task/controllers"
	"backend-onboard-task/middleware"
    "backend-onboard-task/services"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
    userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

    boardService := services.NewBoardService(db)
    boardController := controllers.NewBoardController(boardService)

    router.POST("/signup", userController.RegisterUser)
    router.POST("/login", userController.LoginUser)
      
    router.GET("/boards", boardController.GetBoards)
    router.GET("/boards/:id", boardController.GetBoard)
    
    boardRoutes := router.Group("/").Use(middleware.JWTAuthMiddleware())
    boardRoutes.POST("/boards", boardController.CreateBoard)
    boardRoutes.PUT("/boards/:id", boardController.UpdateBoard)
    boardRoutes.DELETE("/boards/:id", boardController.DeleteBoard)
}