package models

type Sevice string

const (
    NEW Sevice = "NEW"  # 신규 디바이스 
    OLD Sevice = "OLD"  # 구형 디바이스 
)

type DeviceStatus string

const (
    Pending   DeviceStatus = "PENDING"    # 대기 상태 
    Completed DeviceStatus = "COMPLETED"  # 업데이트 완료 상태 
    Failed    DeviceStatus = "FAILED"     # 업데이트 실패 상태 
)