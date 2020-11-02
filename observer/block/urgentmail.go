package main

//urgentMail embedding defaultMail it means that this struct have default mail all filed and function include observer filed
//Attach function Notice function .e.t.c.
type urgentMail struct {
	*defaultMail
}

type info struct {
	to   string
	from string
	body string
}

func NewUrgentMail() *urgentMail {
	return &urgentMail{defaultMail: NewDefaultMail()}
}
