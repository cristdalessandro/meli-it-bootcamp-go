package main

import (
	"errors"
	"fmt"
)

var users []*User

type User struct {
	contract string
	name     string
	DNI      string
	phone    string
	address  string
}

func (user *User) signUp() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Fin de la ejecución")
		if err != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
	}()

	for _, u := range users {
		if u.DNI == user.DNI {
			panic("error: el cliente ya existe")
		}
	}
	_, err := checkUserData(user)
	if err != nil {
		panic(err.Error())
	}
	users = append(users, user)
}

func checkUserData(user *User) (*User, error) {
	if user.contract == "" || user.name == "" || user.DNI == "" || user.phone == "" || user.address == "" {
		return nil, errors.New("validation error")
	}
	return user, nil
}
