package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User는 사용자 인증 정보를 나타내는 애그리거트 루트
type User struct {
	id       string
	pwHash   string // 비밀번호 해시
	isLocked bool   // 계정 잠금 상태
}


// NewUser는 새로운 User를 생성
func NewUser(id, pw string) (*User, error) {
    if id == "" {
        return nil, fmt.Errorf("user ID cannot be empty")
    }
    if len(pw) < 8 {
        return nil, fmt.Errorf("password must be at least 8 characters long")
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
    if err != nil {
        return nil, fmt.Errorf("failed to hash password: %w", err)
    }
    u := &User{
        id:            id,
        pwHash:        string(hash),
        isLocked:      false,
    }

    return u, nil
}

// Authenticate는 사용자 인증을 수행
func (u *User) Authenticate(pw string) bool {
	if u.isLocked {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(u.pwHash), []byte(pw)) == nil
}

func (u *User) Lock() {
    u.isLocked = true
}

func (u *User) GetID() string {
	return u.id
}