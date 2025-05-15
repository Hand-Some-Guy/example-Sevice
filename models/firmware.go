package models

import (
	"fmt"
)

type Sevice string

const (
	NEW Sevice = "NEW"
	OLD Sevice = "OLD"
)

type Firmware struct {
	ID      string `json:"id"`
	Service Sevice `json:"service"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

func NewFirmware(id string, serviceType Sevice, version, path string) (*Firmware, error) {
	if id == "" {
		return nil, fmt.Errorf("firmware ID cannot be empty")
	}
	if serviceType != NEW && serviceType != OLD {
		return nil, fmt.Errorf("invalid service type: %s", serviceType)
	}
	if version == "" {
		return nil, fmt.Errorf("version cannot be empty")
	}
	f := &Firmware{
		ID:      id,
		Service: serviceType,
		Version: version,
		Path:    path,
	}
	return f, nil
}

func (f *Firmware) GetID() string {
	return f.ID
}

func (f *Firmware) GetService() Sevice {
	return f.Service
}

func (f *Firmware) GetVersion() string {
	return f.Version
}

func (f *Firmware) GetPath() string {
	return f.Path
}

func (f *Firmware) ValidateFilePath(path string) error {
	return nil
}