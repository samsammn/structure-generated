package usecase

import (
	"testing"

	"github.com/Komunitas-Mea/marketplace-mea-api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) Find(userID int) (entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.User), args.Error(1)
}

func (mock *MockUserRepository) FindAll() ([]entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.User), args.Error(1)
}

func (mock *MockUserRepository) Create(entity.User) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockUserRepository) Update(entity.User) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockUserRepository) Delete(userID int) error {
	args := mock.Called()
	return args.Error(0)
}

func TestUserFind(t *testing.T) {
	mockRepo := &MockUserRepository{}

	userID := 1
	user := entity.User{ID: userID, Fullname: "Samsam"}

	// Setup expectations
	mockRepo.On("Find").Return(user, nil)

	userUseCase := NewUserUseCase(mockRepo)
	resultData, resultErr := userUseCase.Find(userID)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Nil(t, resultErr)
	assert.Equal(t, userID, resultData.ID)
	assert.Equal(t, "Samsam", resultData.Fullname)
}

func TestUserFindAll(t *testing.T) {
	mockRepo := &MockUserRepository{}

	userID := 1
	user := entity.User{ID: userID, Fullname: "Samsam"}

	// Setup expectations
	mockRepo.On("FindAll").Return([]entity.User{user}, nil)

	userUseCase := NewUserUseCase(mockRepo)
	resultData, resultErr := userUseCase.FindAll()

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Nil(t, resultErr)
	assert.Equal(t, userID, resultData[0].ID)
	assert.Equal(t, "Samsam", resultData[0].Fullname)
}

func TestUserCreate(t *testing.T) {
	mockRepo := &MockUserRepository{}

	userID := 1
	user := entity.User{ID: userID, Fullname: "Samsam", Email: "samsam@mail.com"}

	// Setup expectations
	mockRepo.On("Create").Return(nil)

	userUseCase := NewUserUseCase(mockRepo)
	resultErr := userUseCase.Create(user)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Nil(t, resultErr)
}

func TestUserUpdate(t *testing.T) {
	mockRepo := &MockUserRepository{}

	userID := 1
	user := entity.User{ID: userID, Fullname: "Samsam", Email: "samsam@mail.com"}

	// Setup expectations
	mockRepo.On("Update").Return(nil)

	userUseCase := NewUserUseCase(mockRepo)
	resultErr := userUseCase.Update(user)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Nil(t, resultErr)
}

func TestUserDelete(t *testing.T) {
	mockRepo := &MockUserRepository{}
	userID := 1

	// Setup expectations
	mockRepo.On("Delete").Return(nil)

	userUseCase := NewUserUseCase(mockRepo)
	resultErr := userUseCase.Delete(userID)

	// Mock Assertion
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Nil(t, resultErr)
}
