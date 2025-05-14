package services

import (
    "fmt"
    "instance-20250512-083940/models"
    "instance-20250512-083940/repositories"
)

type FirmwareService interface {
    CreateFirmware(id, serviceType, version, path string) (*models.Firmware, error)
    DeleteFirmware(id string) error
    GetLatestFirmwareByService(serviceType string) (*models.Firmware, error)
}

type firmwareService struct {
    repo repositories.FirmwareRepository
}

func NewFirmwareService(repo repositories.FirmwareRepository) FirmwareService {
    return &firmwareService{repo: repo}
}

func (s *firmwareService) CreateFirmware(id, serviceType, version, path string) (*models.Firmware, error) {
    firmware, err := models.NewFirmware(id, models.Sevice(serviceType), version, path)
    if err != nil {
        return nil, fmt.Errorf("failed to create firmware: %w", err)
    }
    if err := s.repo.Save(firmware); err != nil {
        return nil, fmt.Errorf("failed to save firmware: %w", err)
    }
    return firmware, nil
}

func (s *firmwareService) DeleteFirmware(id string) error {
    if err := s.repo.DeleteByID(id); err != nil {
        return fmt.Errorf("failed to delete firmware: %w", err)
    }
    return nil
}

func (s *firmwareService) GetLatestFirmwareByService(serviceType string) (*models.Firmware, error) {
    firmware, err := s.repo.FindByService(models.Sevice(serviceType))
    if err != nil {
        return nil, fmt.Errorf("failed to find firmware: %w", err)
    }
    return firmware, nil
}