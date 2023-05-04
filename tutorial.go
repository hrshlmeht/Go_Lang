package main

import "fmt"

func main(){

	fmt.Println("Quiz game in go lang")
	var name string 
	fmt.Println("Enter your name : ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v ,welcome to the quiz game\n", name)
	fmt.Printf("Enter your age")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yes , you can play this game!")
	} else {
		fmt.Println("You cannot play this game")
		return
	}

	fmt.Printf("What is better , the RTX 3090 or GTX 3080? ")
	var answer string
	var answer2 string 
	fmt.Scan(&answer , &answer2)
	fmt.Println(answer , answer2)

	if answer + " " + answer2 == "RTX 3090"{
		fmt.Println("Correct")
	}else {
		fmt.Println("Incorrect")
	}
}


