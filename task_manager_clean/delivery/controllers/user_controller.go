package controllers

import (
	"net/http"
	"task_manager/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecaseInterface
}

func NewUserController(usecase domain.UserUsecaseInterface) *UserController {
	return &UserController{
		UserUsecase: usecase,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var user domain.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registeredUser, err := c.UserUsecase.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, registeredUser)
}

func (c *UserController) Login(ctx *gin.Context) {
	var authUser domain.AuthUser
	if err := ctx.BindJSON(&authUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.UserUsecase.LoginUser(&authUser)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
