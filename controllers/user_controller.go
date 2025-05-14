package controllers

import (
	"net/http"
	"instance-20250512-083940/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Login(ctx *gin.Context) {
	var req struct {
		ID       string `json:"id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
		return
	}

	token, err := c.service.Login(req.ID, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "로그인 성공",
		"token":   token,
	})
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	r.POST("/login", c.Login)
}