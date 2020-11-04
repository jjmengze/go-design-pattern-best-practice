package main

import (
	"time"
)

func main() {
	urgentMail := NewUrgentMail()
	userA := NewUser("A-user", urgentMail, make(chan *info), time.Second*2)
	go func() { userA.Observe() }()
	userB := NewUser("B-user", urgentMail, make(chan *info), time.Nanosecond*500)
	go func() { userB.Observe() }()

	userC := NewUser("C-user", urgentMail, make(chan *info), time.Nanosecond)

	urgentMail.Attach(userA.name, userA.info)
	urgentMail.Notify(&info{
		to:   "ALL",
		from: "admin",
		body: " Welcome user A",
	})
	urgentMail.Attach(userB.name, userB.info)
	urgentMail.Notify(&info{
		to:   "ALL",
		from: "admin",
		body: " Welcome  user B",
	})
	urgentMail.Attach(userC.name, userB.info)
	urgentMail.Notify(&info{
		to:   "ALL",
		from: "admin",
		body: " Welcome  user C",
	})
	time.Sleep(time.Second * 30)
}
