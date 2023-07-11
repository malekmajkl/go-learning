package main

import "log"

type User struct {
	FirstName string
	LastName string
}

func main() {
	myMap := make(map[string]User)

	me := User {
		FirstName: "Michal",
		LastName: "Malek",
	}

	myMap["me"] = me
	log.Println(myMap["me"])

	var myNewVar float32

	myNewVar = 11.1

	
}