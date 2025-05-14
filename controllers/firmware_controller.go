package controllers

import (
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
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청입니다"})
        return
    }
    firmware, err := c.service.CreateFirmware(req.ID, req.ServiceType, req.Version, req.Path)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, firmware)
}

func (c *FirmwareController) DeleteFirmware(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := c.service.DeleteFirmware(id); err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "펌웨어가 삭제되었습니다"})
}

func (c *FirmwareController) GetLatestFirmwareByService(ctx *gin.Context) {
    serviceType := ctx.Param("service")
    firmware, err := c.service.GetLatestFirmwareByService(serviceType)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, firmware)
}

func (c *FirmwareController) RegisterRoutes(r *gin.Engine) {
    r.POST("/firmwares", c.CreateFirmware)
    r.DELETE("/firmwares/:id", c.DeleteFirmware)
    r.GET("/firmwares/service/:service", c.GetLatestFirmwareByService)
}