package repositories

import (
	"fmt"
	"instance-20250512-083940/models"
)

type UserRepository interface {
	FindByID(id string) (*models.User, error)
	Save(user *models.User) error
}

type InMemoryUserRepository struct {
	users map[string]*models.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

func (r *InMemoryUserRepository) FindByID(id string) (*models.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found: %s", id)
	}
	return user, nil
}

func (r *InMemoryUserRepository) Save(user *models.User) error {
	r.users[user.GetID()] = user
	return nil
}