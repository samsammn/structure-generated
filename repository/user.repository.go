package repository

import (
	"github.com/Komunitas-Mea/marketplace-mea-api/entity"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

type UserRepository interface {
	// Reader
	Find(userID int) (entity.User, error)
	FindAll() ([]entity.User, error)

	// Writer
	Create(entity.User) error
	Update(entity.User) error
	Delete(userID int) error
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r userRepo) Find(userID int) (user entity.User, err error) {
	query := "SELECT id, fullname, username, phone, email, password FROM users WHERE id=?"
	if err := r.db.Get(&user, query, userID); err != nil {
		return user, err
	}

	return
}

func (r userRepo) FindAll() (users []entity.User, err error) {
	query := "SELECT id, fullname, username, phone, email, password FROM users"
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return
}

func (r userRepo) Create(user entity.User) (err error) {
	_, err = r.db.NamedExec(`INSERT INTO users 
		(fullname, username, email, phone, password) VALUES 
		(:fullname, :username, :email, :phone, :password)`,
		user,
	)

	return
}

func (r userRepo) Update(user entity.User) (err error) {
	_, err = r.db.NamedExec(`UPDATE users SET
		fullname=:fullname, username=:username, email=:email, phone=:phone, password=:password
		WHERE id=:id`,
		user,
	)

	return
}

func (r userRepo) Delete(userID int) (err error) {
	_, err = r.db.Exec(`DELETE FROM users WHERE id=?`,
		userID,
	)

	return
}
