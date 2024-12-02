package main

import "fmt"

const jwtToken string = "abhijeet123" //public

func main() {
	var username string = "Abhijeet"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	var number int = 12
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)
	var precisedNumber float32 = 23.4324543
	fmt.Println(precisedNumber)
	fmt.Printf("Variable is of type: %T \n", precisedNumber)
	var precisedNumber2 float64 = 23.534332
	fmt.Println(precisedNumber2)
	fmt.Printf("Variable is of type: %T \n", precisedNumber2)
	var email string
	fmt.Println(email)
	var emailId = "abhi@go.dev"
	fmt.Println(emailId)
	numofUsers := 12
	fmt.Println(numofUsers)
	var a, b, c int = 10, 20, 30
	fmt.Println(a, b, c)
	var localVar int = 200 //localvariable
	fmt.Println(localVar)
}
