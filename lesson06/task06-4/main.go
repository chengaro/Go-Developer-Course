package main

import "fmt"

type user struct {
	ID   int
	Name string
	contact
}

type employee struct {
	ID   int
	Name string
	contact
}

type contact struct {
	Addss string
	Phone string
}

func main() {
	u := user{
		ID:   1,
		Name: "UserName",
		contact: contact{
			Addss: "UserAddress",
			Phone: "UserPhone",
		},
	}

	e := employee{
		ID:   1,
		Name: "EmployeeName",
		contact: contact{
			Addss: "EmployeeAddress",
			Phone: "EmployeePhone",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
