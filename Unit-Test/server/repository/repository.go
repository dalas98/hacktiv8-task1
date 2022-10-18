package repository

import "dalas98/golang-lesson/Unit-Test/server/model"

type UserRepository interface {
	FindById(id string) *model.User
	CreateNewUser(user *model.User) error
}
