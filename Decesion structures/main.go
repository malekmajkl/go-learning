package main

import "log"


func main() {

	isTrue := false
// IF
	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	myNum := 100
	if myNum > 99 && !isTrue {
		log.Println("mynum is greated than 99 and isTrue is true")
		} else {
		log.Println("mynum is greated than 99 and isTrue is true")
	}

	// switch
	myVar := "dog"
	switch myVar {
	case "cat":
		log.Println("cat is set to cat")

	case "dog":
		log.Println("cat is set to dog")

	case "fish":
		log.Println("cat is set to fish")

	default:
		log.Println("cat is something else")
	}
}
