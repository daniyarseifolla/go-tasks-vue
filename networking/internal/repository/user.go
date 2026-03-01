package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"user-service/internal/model"
)

type UserRepository struct {
	mu    sync.RWMutex
	users map[string]*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*model.User),
	}
}

func (r *UserRepository) Create(name, email string, age int32) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	user := &model.User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Age:       age,
		CreatedAt: now,
		UpdatedAt: now,
	}

	r.users[user.ID] = user
	return user, nil
}

func (r *UserRepository) GetByID(id string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found: %s", id)
	}
	return user, nil
}

func (r *UserRepository) Update(id, name, email string, age int32) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found: %s", id)
	}

	user.Name = name
	user.Email = email
	user.Age = age
	user.UpdatedAt = time.Now()

	return user, nil
}

func (r *UserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return fmt.Errorf("user not found: %s", id)
	}

	delete(r.users, id)
	return nil
}
