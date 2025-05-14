// repositories/firmware_repository.go
package repositories

import (
	"fmt"
    "instance-20250512-083940/models"
)

type FirmwareRepository interface {
	FindByID(id string) (*models.Firmware, error)
	FindByService(service models.Sevice) (*models.Firmware, error)
	Save(f *models.Firmware) error
}

type InMemoryFirmwareRepository struct {
    firmwares map[string]*models.Firmware
}

func NewFirmwareRepository() *InMemoryFirmwareRepository {
    return &InMemoryFirmwareRepository{
		firmwares: make(map[string]*models.Firmware),
	}
}

func (r *InMemoryFirmwareRepository) Save(f *models.Firmware) error {
    r.firmwares[f.GetID()] = f
    return nil
}

// FindByService

func (r *InMemoryFirmwareRepository) FindByID(id string) (*models.Firmware, error) {
	firmware, exists := r.firmwares[id]
	if !exists {
		return nil, fmt.Errorf("firmware not found: %s", id)
	}
	return firmware, nil
}

func (r *InMemoryFirmwareRepository) FindByService(service models.Sevice) (*models.Firmware, error) {
	var latestFirmware *models.Firmware
	var latestVersion string

	for _, firmware := range r.firmwares {
		if firmware.GetService() == service {
			if latestFirmware == nil || compareVersions(firmware.GetVersion(), latestVersion) > 0 {
				latestFirmware = firmware
				latestVersion = firmware.GetVersion()
			}
		}
	}

	if latestFirmware == nil {
		return nil, fmt.Errorf("firmware not found for service: %s", service)
	}

	return latestFirmware, nil
}

// compareVersions compares two version strings (e.g., "1.0.0" vs "1.0.1").
// Returns 1 if v1 > v2, 0 if v1 == v2, -1 if v1 < v2.
func compareVersions(v1, v2 string) int {
	// 간단한 문자열 비교 (실제로는 semver 라이브러리 사용 권장)
	if v1 == v2 {
		return 0
	}
	// 문자열 비교는 정확하지 않으므로, 실제 구현에서는 semver 사용
	if v1 > v2 {
		return 1
	}
	return -1
}