package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUser() []*User {
	return users
}

func AddUser(usr User) (User, error) {
	if usr.Id == 0 {
		return User{}, errors.New("New user must have ID zero")
	}
	usr.Id = nextId
	nextId++

	users = append(users, &usr)

	return usr, nil
}

func GetUserById(id int) (User, error) {
	for _, usr := range users {
		if usr.Id == id {
			return *usr, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' was not found", id)
}

func UpdateUser(usr User) (User, error) {
	for i, currUsr := range users {
		if usr.Id == currUsr.Id {
			users[i] = &usr
			return usr, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' could not be updated", usr.Id)
}

func RemoveUserById(id int) error {
	for i, usr := range users {
		if usr.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' was not found for removal", id)
}
