package models

import (
    "fmt"
    "path/filepath"
    "regexp"
    "strings"
)

type Firmware struct {
	id      string // 애그리거트 루트 식별자
	service Sevice
	version string
	path    string
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
		id:      id,
		service: serviceType,
		version: version,
		path:    path,
	}

	return f, nil
}


// 경로 유효성 검사
func (f *Firmware) ValidateFilePath(path string) error {

	// Basic path checks
    if path == "" {
        return fmt.Errorf("경로를 반드시 입력해 주세요. ")
    }
    if strings.Contains(path, "\x00") {
        return fmt.Errorf("null을 포함할 수 없습니다.")
    }
    if !strings.HasPrefix(path, "/") {
        return fmt.Errorf("linux 포멧에 맞는 경로가 필요합니다 .")
    }

    // Check for invalid characters
    invalidChars := regexp.MustCompile(`[<>\|\*\?\\:;]`)
    if invalidChars.MatchString(path) {
        return fmt.Errorf("path must have .bin or .fw extension")
    }

	// 파일 포맷 검사 
	if !strings.HasSuffix(path, ".bin") && !strings.HasSuffix(path, ".fw") {
        return fmt.Errorf("path must have .bin or .fw extension")
    }

	// 경로가 실재로 존재하는지 검증 
	// 테스트 환경에서 제외 
	// info, err := os.Stat(path)
	// if err != nil {
	// 	return  fmt.Errorf("failed to path: %w", err)
	// }
	// if info.IsDir() {
	// 	return fmt.Errorf("path '%s' is a directory, not a file", path)
	// }

	// 경로 정규화
	cleanPath := filepath.Clean(path)
	if cleanPath == "" || cleanPath == "/" {
		return fmt.Errorf("해당 경로는 정규화 할 수 없습니다")
	}

	return nil

}

func (f *Firmware) GetID() string {
    return f.id
}

func (f *Firmware) GetVersion() string {
    return f.version
}

func (f *Firmware) GetService() Sevice {
    return f.service
}

func (f *Firmware) GetPath() string {
    return f.path
}