package main

func main() {
	urgentMail := NewUrgentMail()
	userA := NewUser("A-user", urgentMail)
	userB := NewUser("B-user", urgentMail)
	//userC := NewUser("C-user", urgentMail)
	urgentMail.Attach(userA.name, userA)
	urgentMail.Notify(&info{
		to:   userA.name,
		from: "admin",
		body: " Welcome user A",
	})
	urgentMail.Attach(userB.name, userB)
	urgentMail.Notify(&info{
		to:   "ALL",
		from: "admin",
		body: " Welcome  user B",
	})
	//urgentMail.Attach(userC.name, userC)
	//urgentMail.Notify(&info{
	//	to:   "ALL",
	//	from: "admin",
	//	body: " Welcome  user C",
	//})

}
