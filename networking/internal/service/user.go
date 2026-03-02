package service

import (
	"fmt"
	"regexp"

	"user-service/internal/model"
	"user-service/internal/repository"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(name, email string, age int32) (*model.User, error) {
	if err := validate(name, email, age); err != nil {
		return nil, err
	}
	return s.repo.Create(name, email, age)
}

func (s *UserService) GetByID(id string) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(id, name, email string, age int32) (*model.User, error) {
	if err := validate(name, email, age); err != nil {
		return nil, err
	}
	return s.repo.Update(id, name, email, age)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}

func validate(name, email string, age int32) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	if age <= 0 {
		return fmt.Errorf("age must be greater than 0")
	}
	return nil
}
