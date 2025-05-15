package controllers

import (
	"log"
	"net/http"
	"instance-20250512-083940/services"

	"github.com/gin-gonic/gin"
)

type FirmwareController struct {
	service services.FirmwareService
}

func NewFirmwareController(service services.FirmwareService) *FirmwareController {
	return &FirmwareController{service: service}
}

func (c *FirmwareController) CreateFirmware(ctx *gin.Context) {
	var req struct {
		ID          string `json:"id" binding:"required"`
		ServiceType string `json:"service_type" binding:"required"`
		Version     string `json:"version" binding:"required"`
		Path        string `json:"path" binding:"required"`
	}
	log.Printf("Received POST /firmwares request from %s: %+v", ctx.ClientIP(), req)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
		return
	}
	firmware, err := c.service.CreateFirmware(req.ID, req.ServiceType, req.Version, req.Path)
	if err != nil {
		log.Printf("Failed to create firmware: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Created firmware: %+v", firmware)
	ctx.JSON(http.StatusCreated, firmware)
}

func (c *FirmwareController) DeleteFirmware(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Printf("Received DELETE /firmwares/%s request from %s", id, ctx.ClientIP())
	if err := c.service.DeleteFirmware(id); err != nil {
		log.Printf("Failed to delete firmware %s: %v", id, err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Deleted firmware %s", id)
	ctx.JSON(http.StatusOK, gin.H{"message": "펌웨어가 삭제되었습니다"})
}

func (c *FirmwareController) GetLatestFirmwareByService(ctx *gin.Context) {
	serviceType := ctx.Param("service")
	log.Printf("Received GET /firmwares/service/%s request from %s", serviceType, ctx.ClientIP())
	firmware, err := c.service.GetLatestFirmwareByService(serviceType)
	if err != nil {
		log.Printf("Failed to find firmware for service %s: %v", serviceType, err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Found firmware for service %s: %+v", serviceType, firmware)
	ctx.JSON(http.StatusOK, firmware)
}

func (c *FirmwareController) RegisterRoutes(r *gin.Engine) {
	r.POST("/firmwares", c.CreateFirmware)
	r.DELETE("/firmwares/:id", c.DeleteFirmware)
	r.GET("/firmwares/service/:service", c.GetLatestFirmwareByService)
}