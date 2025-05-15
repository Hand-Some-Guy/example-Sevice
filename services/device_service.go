package services

import (
	"fmt"
	"instance-20250512-083940/models"
	"instance-20250512-083940/repositories"
)

type DeviceService interface {
	CreateDevice(id, serviceType string) (*models.Device, error)
	GetDevice(id string) (*models.Device, error)
	ChangeDeviceState(id string, status models.DeviceStatus) error
}

type deviceService struct {
	deviceRepo   repositories.DeviceRepository
	firmwareRepo repositories.FirmwareRepository
}

func NewDeviceService(deviceRepo repositories.DeviceRepository, firmwareRepo repositories.FirmwareRepository) DeviceService {
	return &deviceService{
		deviceRepo:   deviceRepo,
		firmwareRepo: firmwareRepo,
	}
}

func (s *deviceService) CreateDevice(id, serviceType string) (*models.Device, error) {
	firmware, err := s.firmwareRepo.FindByService(models.Sevice(serviceType))
	if err != nil {
		return nil, fmt.Errorf("failed to find firmware: %w", err)
	}
	device, err := models.NewDevice(id, serviceType, firmware.GetID(), firmware.GetVersion())
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %w", err)
	}
	if err := s.deviceRepo.Save(device); err != nil {
		return nil, fmt.Errorf("failed to save device: %w", err)
	}
	return device, nil
}

func (s *deviceService) GetDevice(id string) (*models.Device, error) {
	device, err := s.deviceRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get device: %w", err)
	}
	return device, nil
}

func (s *deviceService) ChangeDeviceState(id string, status models.DeviceStatus) error {
	device, err := s.deviceRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to get device: %w", err)
	}
	if err := device.SetStatus(status); err != nil {
		return fmt.Errorf("failed to change state: %w", err)
	}
	if err := s.deviceRepo.Save(device); err != nil {
		return fmt.Errorf("failed to save device: %w", err)
	}
	return nil
}