package services

import (
	"fmt"
	"time"
	"instance-20250512-083940/repositories"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	Login(id, password string) (string, error)
}

type userService struct {
	repo       repositories.UserRepository
	jwtSecret  string
}

func NewUserService(repo repositories.UserRepository, jwtSecret string) UserService {
	return &userService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *userService) Login(id, password string) (string, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return "", fmt.Errorf("로그인 실패: %w", err)
	}
	if !user.Authenticate(password) {
		return "", fmt.Errorf("잘못된 비밀번호 또는 계정이 잠겨 있습니다")
	}

	// JWT 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("토큰 생성 실패: %w", err)
	}

	return tokenString, nil
}