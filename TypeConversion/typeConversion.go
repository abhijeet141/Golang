package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter an integer")
	number, _ := reader.ReadString('\n')
	number = strings.TrimSpace(number)
	fmt.Println(number)
	fmt.Printf("%T \n", number)
	//string to integer
	intValue, err1 := strconv.Atoi(number)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(intValue)
	fmt.Printf("%T \n", intValue)
	//string to float
	floatValue, err2 := strconv.ParseFloat(number, 64)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Printf("%.2f \n", floatValue)
	fmt.Printf("%T \n", floatValue)
	result := float64(intValue) + floatValue
	fmt.Println(result)
}
