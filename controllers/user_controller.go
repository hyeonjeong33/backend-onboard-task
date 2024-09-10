package controllers

import (
    "net/http"
	"github.com/gin-gonic/gin"
	"backend-onboard-task/helpers"
    "backend-onboard-task/models"
    "backend-onboard-task/services"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
	var input models.User

	if !helpers.BindJSON(ctx, &input) {
		return
	}
	
	if isValid, message := helpers.ValidateUserInput(input.Email, input.Password); !isValid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	if err := c.UserService.RegisterUser(input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "회원가입이 성공적으로 완료되었습니다."})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var input models.User

	if !helpers.BindJSON(ctx, &input) {
		return
	}
	
	if isValid, message := helpers.ValidateUserInput(input.Email, input.Password); !isValid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	token, err := c.UserService.LoginUser(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "로그인에 성공했습니다.", "token": token})
}