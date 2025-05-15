package models

import (
	"fmt"
)

type Device struct {
	ID              string       `json:"id"`
	Service         Sevice       `json:"service"`
	FirmwareID      string       `json:"firmwareID"`
	FirmwareVersion string       `json:"firmwareVersion"`
	Status          DeviceStatus `json:"status"`
}

func NewDevice(id, serviceType, firmwareID, firmwareVersion string) (*Device, error) {
	if id == "" {
		return nil, fmt.Errorf("device ID cannot be empty")
	}
	if serviceType != string(NEW) && serviceType != string(OLD) {
		return nil, fmt.Errorf("invalid service type: %s", serviceType)
	}
	if firmwareID == "" {
		return nil, fmt.Errorf("firmware ID cannot be empty")
	}
	if firmwareVersion == "" {
		return nil, fmt.Errorf("firmware version cannot be empty")
	}
	return &Device{
		ID:              id,
		Service:         Sevice(serviceType),
		FirmwareID:      firmwareID,
		FirmwareVersion: firmwareVersion,
		Status:          Pending,
	}, nil
}

func (d *Device) SetStatus(status DeviceStatus) error {
	if status != Pending && status != Completed && status != Failed {
		return fmt.Errorf("invalid status: %s", status)
	}
	d.Status = status
	return nil
}

func (d *Device) GetID() string {
	return d.ID
}

func (d *Device) GetService() Sevice {
	return d.Service
}

func (d *Device) GetFirmwareID() string {
	return d.FirmwareID
}

func (d *Device) GetFirmwareVersion() string {
	return d.FirmwareVersion
}

func (d *Device) GetStatus() DeviceStatus {
	return d.Status
}