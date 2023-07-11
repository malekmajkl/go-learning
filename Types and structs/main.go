package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}
q
func main() {
	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
		PhoneNumber: "890 90 89 67",
	}

	log.Println("first name is", user.FirstName, "Last name is", user.LastName)
}
