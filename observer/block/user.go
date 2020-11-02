package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
	mail *urgentMail
}

func NewUser(name string, mail *urgentMail) *user {
	return &user{
		name: name,
		mail: mail,
	}
}

func (u user) Update(body *info) error {
	fmt.Println(u.name + "get a mail")
	fmt.Println(body)
	// do something in a long time
	// it will block other user
	time.Sleep(time.Second * 5)
	return nil
}
