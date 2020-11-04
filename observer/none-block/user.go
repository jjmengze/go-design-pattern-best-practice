package main

import (
	"fmt"
	"time"
)

type user struct {
	name     string
	mail     *urgentMail
	overload time.Duration
	info     chan *info
}

func NewUser(name string, mail *urgentMail, info chan *info, overTime time.Duration) *user {
	return &user{
		name:     name,
		mail:     mail,
		info:     info,
		overload: overTime,
	}
}

func (u user) Observe() error {
	for info := range u.info {
		fmt.Println(u.name + " get a mail")
		fmt.Println(info.body)
		time.Sleep(u.overload)
	}

	// do something in a long time
	// it will block other user
	return nil
}
