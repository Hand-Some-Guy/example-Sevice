package controllers

import (
	"log"
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
	log.Printf("Received POST /devices request from %s: ID=%s, ServiceType=%s",
		ctx.ClientIP(), req.ID, req.ServiceType)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
		return
	}
	device, err := c.service.CreateDevice(req.ID, req.ServiceType)
	if err != nil {
		log.Printf("Failed to create device: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Created device: ID=%s, Service=%s, FirmwareID=%s, Status=%s",
		device.GetID(), device.GetService(), device.GetFirmwareID(), device.GetStatus())
	ctx.JSON(http.StatusCreated, device)
}

func (c *DeviceController) GetDevice(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Printf("Received GET /devices/%s request from %s", id, ctx.ClientIP())
	device, err := c.service.GetDevice(id)
	if err != nil {
		log.Printf("Failed to get device %s: %v", id, err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Found device: ID=%s, Service=%s, FirmwareID=%s, Status=%s",
		device.GetID(), device.GetService(), device.GetFirmwareID(), device.GetStatus())
	ctx.JSON(http.StatusOK, device)
}

func (c *DeviceController) ChangeDeviceState(ctx *gin.Context) {
	id := ctx.Param("id")
	var req struct {
		State string `json:"state" binding:"required"`
	}
	log.Printf("Received PUT /devices/%s/state request from %s: State=%s",
		id, ctx.ClientIP(), req.State)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
		return
	}
	if err := c.service.ChangeDeviceState(id, models.DeviceStatus(req.State)); err != nil {
		log.Printf("Failed to change device state for %s: %v", id, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Changed device %s state to %s", id, req.State)
	ctx.JSON(http.StatusOK, gin.H{"message": "state changed"})
}

func (c *DeviceController) RegisterRoutes(r *gin.Engine) {
	r.POST("/devices", c.CreateDevice)
	r.GET("/devices/:id", c.GetDevice)
	r.PUT("/devices/:id/state", c.ChangeDeviceState)
}