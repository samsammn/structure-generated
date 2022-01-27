package usecase

import (
	"github.com/Komunitas-Mea/marketplace-mea-api/entity"
	"github.com/Komunitas-Mea/marketplace-mea-api/repository"
)

type userUseCase struct {
	repo repository.UserRepository
}

type UserUseCase interface {
	// Reader
	Find(userID int) (entity.User, error)
	FindAll() ([]entity.User, error)

	// Writer
	Create(entity.User) error
	Update(entity.User) error
	Delete(userID int) error
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (uc userUseCase) Find(userID int) (entity.User, error) {
	return uc.repo.Find(userID)
}

func (uc userUseCase) FindAll() ([]entity.User, error) {
	return uc.repo.FindAll()
}

func (uc userUseCase) Create(user entity.User) error {
	return uc.repo.Create(user)
}

func (uc userUseCase) Update(user entity.User) error {
	return uc.repo.Update(user)
}

func (uc userUseCase) Delete(userID int) error {
	return uc.repo.Delete(userID)
}
