// Code generated by mockery v2.9.4. DO NOT EDIT.

package repositoryuser

import (
	entity "dalas98/golang-lesson/Final-Project/entity"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryUserMock is an autogenerated mock type for the RepositoryUser type
type RepositoryUserMock struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *RepositoryUserMock) Create(data entity.User) (entity.User, error) {
	ret := _m.Called(data)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: id
func (_m *RepositoryUserMock) DeleteByID(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsEmailExist provides a mock function with given fields: email
func (_m *RepositoryUserMock) IsEmailExist(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: email
func (_m *RepositoryUserMock) Login(email string) (entity.User, error) {
	ret := _m.Called(email)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: data
func (_m *RepositoryUserMock) Update(data entity.User) (entity.User, error) {
	ret := _m.Called(data)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
