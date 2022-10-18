package service

import (
	"dalas98/golang-lesson/Unit-Test/server/model"
	"dalas98/golang-lesson/Unit-Test/server/repository"
	"database/sql"
)

type UserServices struct {
	repo repository.UserRepository
}

func NewUserServices(repo repository.UserRepository) *UserServices {
	return &UserServices{
		repo: repo,
	}
}

func (u *UserServices) GetSingleUserById(id string) (*model.User, error) {
	user := u.repo.FindById(id)
	if user == nil {
		return nil, sql.ErrNoRows
	}

	product := u.repo.FindById(id)
	if product == nil {
		return nil, sql.ErrNoRows
	}

	return user, nil
}

// func (u *UserServices) CreateNewUser(user *model.User)
