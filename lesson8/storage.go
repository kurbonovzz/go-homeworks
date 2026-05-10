package main

import "fmt"

type Storage struct {
	users  map[int64]User
	nextID int64
}

func NewStorage() *Storage {
	return &Storage{
		users:  make(map[int64]User),
		nextID: 1,
	}
}

func (s *Storage) CreateUser(name string, age int) User {

	user := User{
		ID:   s.nextID,
		Name: name,
		Age:  age,
	}

	s.users[user.ID] = user
	s.nextID = s.nextID + 1

	return user
}

func (s *Storage) GetUser(id int64) {
	user, ok := s.users[id]

	if ok {
		fmt.Println(user)
	} else {
		fmt.Println("не найден")
	}
}

func (s *Storage) UpdateUser(id int64, name string, age int) {

	user, ok := s.users[id]

	if !ok {
		fmt.Println("не найден")
		return
	}

	user.Name = name
	user.Age = age

	s.users[id] = user
}

func (s *Storage) DeleteUser(id int64) {
	delete(s.users, id)
}
