package memory

import (
	"github.com/gianlucas34/ecommerce-api/internal/domain"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type UserRepository struct {
	Users             []*domain.User
	createFailed      bool
	findAllFailed     bool
	findFailed        bool
	findByEmailFailed bool
	emailAlreadyInUse bool
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Users: make([]*domain.User, 0),
	}
}

func (r *UserRepository) Create(user *domain.User) error {
	if r.createFailed {
		return errors.NewInternalServerError(errors.CREATE_USER)
	}

	r.Users = append(r.Users, user)

	return nil
}

func (r *UserRepository) FindAll() ([]*domain.User, error) {
	if r.findAllFailed {
		return nil, errors.NewInternalServerError(errors.FIND_ALL_USERS)
	}

	return r.Users, nil
}

func (r *UserRepository) Find(id string) (*domain.User, error) {
	if r.findFailed {
		return nil, errors.NewInternalServerError(errors.FIND_USER)
	}

	for _, user := range r.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (bool, error) {
	if r.findByEmailFailed {
		return false, errors.NewInternalServerError(errors.FIND_USER_BY_EMAIL)
	}

	if r.emailAlreadyInUse {
		return true, nil
	}

	for _, user := range r.Users {
		if user.Email == email {
			return true, nil
		}
	}

	return false, nil
}

func (r *UserRepository) FailCreate() *UserRepository {
	r.createFailed = true

	return r
}

func (r *UserRepository) FailFindAll() *UserRepository {
	r.findAllFailed = true

	return r
}

func (r *UserRepository) FailFind() *UserRepository {
	r.findFailed = true

	return r
}

func (r *UserRepository) FailFindByEmail() *UserRepository {
	r.findByEmailFailed = true

	return r
}

func (r *UserRepository) EmailAlreadyInUse() *UserRepository {
	r.emailAlreadyInUse = true

	return r
}
