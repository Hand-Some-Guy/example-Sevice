package repositories

import (
    "fmt"
    "instance-20250512-083940/models"
    "github.com/Masterminds/semver/v3"
)

type FirmwareRepository interface {
    FindByID(id string) (*models.Firmware, error)
    FindByService(service models.Sevice) (*models.Firmware, error)
    Save(f *models.Firmware) error
    DeleteByID(id string) error
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

func (r *InMemoryFirmwareRepository) FindByID(id string) (*models.Firmware, error) {
    firmware, exists := r.firmwares[id]
    if !exists {
        return nil, fmt.Errorf("firmware not found: %s", id)
    }
    return firmware, nil
}

func (r *InMemoryFirmwareRepository) FindByService(service models.Sevice) (*models.Firmware, error) {
    var latestFirmware *models.Firmware
    var latestVersion *semver.Version

    for _, firmware := range r.firmwares {
        if firmware.GetService() == service {
            currentVersion, err := semver.NewVersion(firmware.GetVersion())
            if err != nil {
                continue // 잘못된 버전은 무시
            }
            if latestFirmware == nil || currentVersion.GreaterThan(latestVersion) {
                latestFirmware = firmware
                latestVersion = currentVersion
            }
        }
    }

    if latestFirmware == nil {
        return nil, fmt.Errorf("firmware not found for service: %s", service)
    }
    return latestFirmware, nil
}

func (r *InMemoryFirmwareRepository) DeleteByID(id string) error {
    if _, exists := r.firmwares[id]; !exists {
        return fmt.Errorf("firmware not found: %s", id)
    }
    delete(r.firmwares, id)
    return nil
}