package main

import (
	"os"
	"fmt"

	"instance-20250512-083940/controllers"
	"instance-20250512-083940/models"
	"instance-20250512-083940/repositories"
	"instance-20250512-083940/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// device 레이어 
	deviceRepo := repositories.NewInMemoryDeviceRepository()
	firmwareRepo := repositories.NewFirmwareRepository()
	service := services.NewDeviceService(deviceRepo, firmwareRepo)
	controller := controllers.NewDeviceController(service)

	// user 레이어 
	userRepo := repositories.NewInMemoryUserRepository()
	userService := services.NewUserService(userRepo, "my-secret-key") // 실제로는 환경 변수 사용 권장
	userController := controllers.NewUserController(userService)

	// firmware 레이어 
    firmwareService := services.NewFirmwareService(firmwareRepo)
    firmwareController := controllers.NewFirmwareController(firmwareService)


	// init Data
	if err := initData(firmwareRepo, userRepo); err != nil {
    	panic(err)
	}

	// router add
	controller.RegisterRoutes(r)
	userController.RegisterRoutes(r)
    firmwareController.RegisterRoutes(r)

	// listen port set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
	}

// 초기 데이터 선언 
func initData(firmwareRepo repositories.FirmwareRepository, userRepo repositories.UserRepository) error {
    firmwares := []struct {
        id      string
        service models.Sevice
        version string
        path    string
    }{
        {"fw1", models.NEW, "1.0.0", "/firmware.bin"},
        {"fw2", models.NEW, "1.0.1", "/firmware_new.bin"},
        {"fw3", models.OLD, "2.0.0", "/firmware_old.bin"},
    }
    for _, f := range firmwares {
        fw, err := models.NewFirmware(f.id, f.service, f.version, f.path)
        if err != nil {
            return fmt.Errorf("failed to create firmware %s: %w", f.id, err)
        }
        if err := firmwareRepo.Save(fw); err != nil {
            return fmt.Errorf("failed to save firmware %s: %w", f.id, err)
        }
    }

    user, err := models.NewUser("testuser", "password123")
    if err != nil {
        return fmt.Errorf("failed to create test user: %w", err)
    }
    if err := userRepo.Save(user); err != nil {
        return fmt.Errorf("failed to save test user: %w", err)
    }
    return nil
}