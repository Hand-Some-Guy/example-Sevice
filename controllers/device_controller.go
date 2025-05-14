package controllers

import (
	"net/http"
	"instance-20250512-083940/models"
	"instance-20250512-083940/services"

	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	service services.DeviceService
}

func NewDeviceController(service services.DeviceService) *DeviceController {
	return &DeviceController{service: service}
}

func (c *DeviceController) CreateDevice(ctx *gin.Context) {
	var req struct {
		ID          string `json:"id" binding:"required"`
		ServiceType string `json:"service_type" binding:"required"`
	}

	// 응답 형식 객체 바인드 
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// 
	device, err := c.service.CreateDevice(req.ID, req.ServiceType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, device)
}

func (c *DeviceController) GetDevice(ctx *gin.Context) {
	id := ctx.Param("id")
	device, err := c.service.GetDevice(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, device)
}

func (c *DeviceController) ChangeDeviceState(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		State string `json:"state" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := c.service.ChangeDeviceState(id, models.DeviceStatus(req.State)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "state changed"})
}

func (c *DeviceController) RegisterRoutes(r *gin.Engine) {
	r.POST("/devices", c.CreateDevice)
	r.GET("/devices/:id", c.GetDevice)
	r.PUT("/devices/:id/state", c.ChangeDeviceState)
}