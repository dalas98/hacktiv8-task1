package postgres

import (
	"dalas98/golang-lesson/Unit-Test/server/model"
	"dalas98/golang-lesson/Unit-Test/server/repository"
	"fmt"
)

type userRepo struct{}

func NewUserRepo() repository.UserRepository {
	return &userRepo{}
}

func (u *userRepo) FindById(id string) *model.User {
	query := "SELECT * FROM users WHERE id=$1"

	fmt.Println(query)
	return nil
}

func (u *userRepo) CreateNewUser(user *model.User) error {
	return nil
}
