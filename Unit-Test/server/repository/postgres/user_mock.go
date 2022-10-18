package postgres

import (
	"dalas98/golang-lesson/Unit-Test/server/model"
	"dalas98/golang-lesson/Unit-Test/server/repository"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	Mock *mock.Mock
}

func NewUserRepoMock(mock *mock.Mock) repository.UserRepository {
	return &UserRepoMock{
		Mock: mock,
	}
}

func (u *UserRepoMock) FindById(id string) *model.User {
	arguments := u.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	user := arguments.Get(0).(model.User)
	return &user
}

func (u *UserRepoMock) CreateNewUser(user *model.User) error {
	arguments := u.Mock.Called(user)

	if arguments.Get(0) == nil {
		return nil
	}

	//  := arguments.Get(0).(model.User)
	return nil
}
