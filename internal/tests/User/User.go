package user_test

import (
	"log"
	"todoapp/internal/entities"
)

type MockUserService struct{}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}
func (m *MockUserService) FindByUser(user entities.User) (*entities.User, error) {

	log.Printf("FindByUser metoduna gelen User: %v", user)
	return nil, nil
}
