package main

import "log"



func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	// animals := []string{"dog", "fish", "horse", "cow"}
	animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "charlie"
	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}

	var firstLine = "once upon a time"
	for i,l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		Firstname string
		LastName string
		Email string
		Age int
	}
	var users []User
	users = append(users, User{"John", "Smith", "johnsmith@gmail.com", 89})
	users = append(users, User{"John", "Smith", "johnsmith@gmail.com", 84})
	users = append(users, User{"John", "Smith", "johnsmith@gmail.com", 81})

	for _, l :=range users {
		log.Println(l.Firstname, l.LastName, l.Email, l.Age)
	}
}
