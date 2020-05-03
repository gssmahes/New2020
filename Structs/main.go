package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	gssm := person{
		lastName:  "Gurusamy",
		firstName: "Maheswaran",
		contactInfo: contactInfo{
			email:   "maheshwaran.gss@gmail.com",
			zipCode: 624301,
		},
	}
	fmt.Println(gssm)
	fmt.Printf("%+v", gssm)
}
