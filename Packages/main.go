package main

import (
	"log"

	"github.com/malek.majkl/go-learning/helpers"
)

func main()  {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)

}

