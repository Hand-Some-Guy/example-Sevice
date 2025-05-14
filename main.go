package main

import (
	"instance-20250512-083940/controllers"
	"instance-20250512-083940/repositories"
	"instance-20250512-083940/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 의존성 조립
	deviceRepo := repositories.NewInMemoryDeviceRepository()
	firmwareRepo := repositories.NewFirmwareRepository()

	service := services.NewDeviceService(deviceRepo, firmwareRepo)
	controller := controllers.NewDeviceController(service)

	// 라우트 등록
	controller.RegisterRoutes(r)

	// 서버 실행
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}