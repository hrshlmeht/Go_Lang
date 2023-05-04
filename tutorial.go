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
	var  score int 
	var answer3 string 
	score = 0
	fmt.Scan(&answer , &answer2)
	fmt.Println(answer , answer2)

	if answer + " " + answer2 == "RTX 3090"{
		fmt.Println("Correct")
		score++
	} else if answer + " " + answer2 == "rtx 3090"{
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("Which university is better , Cornell or Gtech?")
	fmt.Scan(&answer3)

	if answer3 == "Cornell"{
		fmt.Printf("Thats correct")
		score++
	} else {
		fmt.Printf("Oops you have answered the question wrong")
	}

	fmt.Printf("The total score of  %v  is %v \n", name, score)

}//main ends


