package repositories

import (
	"fmt"
	"instance-20250512-083940/models"
)

type DeviceRepository interface {
	Save(device *models.Device) error
	FindByID(id string) (*models.Device, error)
}

type InMemoryDeviceRepository struct {
	devices map[string]*models.Device
}

func NewInMemoryDeviceRepository() *InMemoryDeviceRepository {
	return &InMemoryDeviceRepository{
		devices: make(map[string]*models.Device),
	}
}

func (r *InMemoryDeviceRepository) Save(device *models.Device) error {
	r.devices[device.GetID()] = device
	return nil
}

func (r *InMemoryDeviceRepository) FindByID(id string) (*models.Device, error) {
	device, exists := r.devices[id]
	if !exists {
		return nil, fmt.Errorf("device not found: %s", id)
	}
	return device, nil
}