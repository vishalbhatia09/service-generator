package main

import (
	"fmt"
)

type Service interface {
	AddUser(int, string) Service
	RemoveUser(int) Service
	SubscribeUser(int) Service
	UnsubscribeUser(int) Service
	Print()
}

type service struct {
	name              string
	usersinfo         []User
	subscriberedUsers []User
}

type User struct {
	name string
	id   int
}

func NewService(name string) Service {
	users := make([]User, 0)
	subscribers := make([]User, 0)
	return &service{
		name:              name,
		usersinfo:         users,
		subscriberedUsers: subscribers,
	}
}

func (s *service) AddUser(id int, name string) Service {
	fmt.Printf("adding user %d to service %s\n", id, s.name)
	tempuser := User{name: name, id: id}

	s.usersinfo = append(s.usersinfo, tempuser)
	s.subscriberedUsers = append(s.subscriberedUsers, tempuser)

	return &service{
		name:              s.name,
		usersinfo:         s.usersinfo,
		subscriberedUsers: s.subscriberedUsers,
	}

}

func (s *service) RemoveUser(id int) Service {
	fmt.Printf("removing user %d from service %s\n", id, s.name)
	var desired_index int
	for i, val := range s.usersinfo {
		if val.id == id {
			desired_index = i
		}
	}

	s.usersinfo = append(s.usersinfo[:desired_index], s.usersinfo[desired_index+1:]...)
	s.subscriberedUsers = append(s.subscriberedUsers[:desired_index], s.subscriberedUsers[desired_index+1:]...)

	return s

}

func (s *service) SubscribeUser(id int) Service {
	fmt.Printf("subscribing user %d to service %s\n", id, s.name)

	s.subscriberedUsers = append(s.subscriberedUsers, User{id: id})

	return s
}

func (s *service) UnsubscribeUser(id int) Service {
	fmt.Printf("unsubscribing user %d from service %s\n", id, s.name)
	var desired_index int
	for i, val := range s.usersinfo {
		if val.id == id {
			desired_index = i
		}
	}

	s.subscriberedUsers = append(s.subscriberedUsers[:desired_index], s.subscriberedUsers[desired_index+1:]...)

	return s

}

func (s *service) Print() {
	fmt.Println("\nservice ", s.name)
	fmt.Println("Users:")
	for _, val := range s.usersinfo {
		fmt.Printf("Id:%d, name:%s\n", val.id, val.name)
	}

	fmt.Println("Subscribers:")
	for _, val := range s.subscriberedUsers {
		fmt.Printf("%d\n", val.id)
	}
}

func main() {
	fmt.Println("Starting...")

	ss := make([]Service, 0)
	s1 := NewService("s1")
	ss = append(ss, s1)

	s1.AddUser(1, "abc")
	s1.AddUser(2, "def")
	s1.AddUser(3, "efg")
	s1.Print()

	s2 := NewService("s2")
	s2.AddUser(4, "wer")
	s2.AddUser(5, "rty")
	s2.AddUser(6, "uio")
	s2.Print()

	s1.RemoveUser(2)
	s1.SubscribeUser(5)
	s1.UnsubscribeUser(3)
	s1.Print()

}
