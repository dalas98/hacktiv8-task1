package main

import (
	"fmt"
	"sync"
)

type Users struct {
	Name string
	Age  int
}

type UserService struct {
	Users []*Users
}

type UsersInterface interface {
	Register(u *Users) string
	GetUsers() []*Users
}

func main() {
	var userIface UserService

	users := []*Users{
		{
			Name: "Safira Angel",
			Age:  26,
		},
		{
			Name: "Clara Monica",
			Age:  22,
		},
		{
			Name: "Fahmi",
			Age:  26,
		},
		{
			Name: "Giva",
			Age:  24,
		},
	}

	var wait sync.WaitGroup
	lenArr := len(users)
	fmt.Println("panjang array", lenArr)
	wait.Add(lenArr)
	for _, v := range users {
		go wrapper(&wait, userIface, v)
	}
	wait.Wait()

	resUser := userIface.GetUsers()
	fmt.Println("Get Users")
	for _, v := range resUser {
		fmt.Printf("Name: %s Age: %d\n", v.Name, v.Age)
	}
}

func NewUser(u []*Users) UsersInterface {
	return &UserService{
		Users: u,
	}
}

func (us *UserService) Register(u *Users) string {
	us.Users = append(us.Users, u)

	return "User " + u.Name + " Berhasil di daftarkan"
}

func (us *UserService) GetUsers() []*Users {
	return us.Users
}

func wrapper(wg *sync.WaitGroup, service UserService, user *Users) {
	res := service.Register(user)
	fmt.Println(res)
	wg.Done()
}
