package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printMenu() {
	fmt.Println("CLI CALCULATOR")
	fmt.Println("1->ADD")
	fmt.Println("2->SUBTRACT")
	fmt.Println("3->MULTIPLY")
	fmt.Println("4->DIVIDE")
	fmt.Println("5->CLEAR")
	fmt.Println("6->EXIT")
}
func add(a, b float64) float64 {
	return a + b
}
func subtract(a, b float64) float64 {
	return a - b
}
func multiply(a, b float64) float64 {
	return a * b
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by 0")
	}
	return a / b, nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		inputChoice, err1 := strconv.Atoi(choice)
		if err1 != nil {
			fmt.Println("Error", err1)
			return
		}
		if inputChoice < 1 || inputChoice > 6 {
			fmt.Println("Enter choice between 1 and 6")
			continue
		}
		if inputChoice == 6 {
			fmt.Println("Exiting")
			break
		}
		if inputChoice == 5 {
			//will work
			fmt.Println("Cleared History")
			continue
		}
		fmt.Print("Enter 1st Number: ")
		num1_str, _ := reader.ReadString('\n')
		num1_str = strings.TrimSpace(num1_str)
		num1, err2 := strconv.ParseFloat(num1_str, 64)
		if err2 != nil {
			fmt.Println("Error", err2)
			return
		}
		fmt.Print("Enter 2nd Number: ")
		num2_str, _ := reader.ReadString('\n')
		num2_str = strings.TrimSpace(num2_str)
		num2, err2 := strconv.ParseFloat(num2_str, 64)
		if err2 != nil {
			fmt.Println("Error", err2)
			return
		}
		var result float64
		var err3 error
		switch inputChoice {
		case 1:
			result = add(num1, num2)
		case 2:
			result = subtract(num1, num2)
		case 3:
			result = multiply(num1, num2)
		case 4:
			result, err3 = divide(num1, num2)
		default:
			fmt.Println("Invalid Choice")
		}
		if err3 != nil {
			fmt.Println(err3)
			continue
		}
		fmt.Printf("Result: %v \n", result)
	}
}
