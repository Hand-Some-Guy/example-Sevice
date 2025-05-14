// repositories/firmware_repository.go
package repositories

import (
	"fmt"
    "instance-20250512-083940/models"
)

type FirmwareRepository interface {
	FindByID(id string) (*models.Firmware, error)
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