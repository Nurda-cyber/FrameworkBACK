package mocks

import (
	"WelcomeGo/models"

	"github.com/stretchr/testify/mock"
)

type MockToyRepository struct {
	mock.Mock
}

func (m *MockToyRepository) GetAll() ([]models.Toy, error) {
	args := m.Called()
	return args.Get(0).([]models.Toy), args.Error(1)
}

func (m *MockToyRepository) GetByID(id uint) (models.Toy, error) {
	args := m.Called(id)
	return args.Get(0).(models.Toy), args.Error(1)
}

func (m *MockToyRepository) Create(toy models.Toy) (models.Toy, error) {
	args := m.Called(toy)
	return args.Get(0).(models.Toy), args.Error(1)
}

func (m *MockToyRepository) Update(id uint, toy models.Toy) (models.Toy, error) {
	args := m.Called(id, toy)
	return args.Get(0).(models.Toy), args.Error(1)
}

func (m *MockToyRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
