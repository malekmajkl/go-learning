package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json: "hair_color"`
	HasDog bool `json:"has_dog"`
}

func main()  {
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wauyne",
		"hair_color": "black",
		"has_dog": false
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err !=nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write JSON from struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "wally"
	m1.LastName = "West"
	m1.HairColor = "yellow"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}