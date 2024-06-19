package memoryrepos

import (
	"github.com/gianlucas34/ecommerce-api/internal/domain"
	"github.com/gianlucas34/ecommerce-api/internal/errors"
)

type UserRepositoryMemory struct {
	Users             []*domain.User
	createFailed      bool
	findAllFailed     bool
	findFailed        bool
	findByEmailFailed bool
	emailAlreadyInUse bool
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	return &UserRepositoryMemory{
		Users: make([]*domain.User, 0),
	}
}

func (r *UserRepositoryMemory) Create(user *domain.User) error {
	if r.createFailed {
		return errors.NewInternalServerError(errors.CREATE_USER)
	}

	r.Users = append(r.Users, user)

	return nil
}

func (r *UserRepositoryMemory) FindAll() ([]*domain.User, error) {
	if r.findAllFailed {
		return nil, errors.NewInternalServerError(errors.FIND_ALL_USERS)
	}

	return r.Users, nil
}

func (r *UserRepositoryMemory) Find(id string) (*domain.User, error) {
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

func (r *UserRepositoryMemory) FindByEmail(email string) (bool, error) {
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

func (r *UserRepositoryMemory) FailCreate() *UserRepositoryMemory {
	r.createFailed = true

	return r
}

func (r *UserRepositoryMemory) FailFindAll() *UserRepositoryMemory {
	r.findAllFailed = true

	return r
}

func (r *UserRepositoryMemory) FailFind() *UserRepositoryMemory {
	r.findFailed = true

	return r
}

func (r *UserRepositoryMemory) FailFindByEmail() *UserRepositoryMemory {
	r.findByEmailFailed = true

	return r
}

func (r *UserRepositoryMemory) EmailAlreadyInUse() *UserRepositoryMemory {
	r.emailAlreadyInUse = true

	return r
}
