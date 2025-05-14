package models

type Sevice string

const (
    NEW Sevice = "NEW"
    OLD Sevice = "OLD"
)

type DeviceStatus string

const (
    Pending   DeviceStatus = "PENDING"
    Completed DeviceStatus = "COMPLETED"
    Failed    DeviceStatus = "FAILED"
)