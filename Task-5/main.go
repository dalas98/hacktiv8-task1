package main

import "fmt"

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
	}
	for _, v := range users {
		resp := userIface.Register(v)
		fmt.Println(resp)
	}

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
