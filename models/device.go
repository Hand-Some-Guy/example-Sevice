package models

import (
    "fmt"
)



type Device struct {
    id              string
    firmwareID      string
    firmwareVersion string
    status          DeviceStatus
}

func NewDevice(id, firmwareID, firmwareVersion string) (*Device, error) {
    if id == "" {
        return nil, fmt.Errorf("device ID cannot be empty")
    }
    if firmwareID == "" {
        return nil, fmt.Errorf("firmware ID cannot be empty")
    }
    if firmwareVersion == "" {
        return nil, fmt.Errorf("firmware version cannot be empty")
    }
    return &Device{
        id:              id,
        firmwareID:      firmwareID,
        firmwareVersion: firmwareVersion,
        status:          Pending,
    }, nil
}

func (d *Device) SetStatus(status DeviceStatus) error {
    if status != Pending && status != Completed && status != Failed {
        return fmt.Errorf("invalid status: %s", status)
    }
    d.status = status
    return nil
}

func (d *Device) GetID() string {
    return d.id
}

func (d *Device) GetFirmwareID() string {
    return d.firmwareID
}

func (d *Device) GetFirmwareVersion() string {
    return d.firmwareVersion
}

func (d *Device) GetStatus() DeviceStatus {
    return d.status
}