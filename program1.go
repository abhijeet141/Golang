package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter first number: ")
	num1str, _ := reader.ReadString('\n')
	num1str = strings.TrimSpace(num1str)
	fmt.Println("Enter second number: ")
	num2str, _ := reader.ReadString('\n')
	num2str = strings.TrimSpace(num2str)
	num1, err1 := strconv.Atoi(num1str)
	if err1 != nil {
		fmt.Println("Error", err1)
		return
	}
	num2, err2 := strconv.Atoi(num2str)
	if err2 != nil {
		fmt.Println("Error", err2)
		return
	}
	result := add(num1, num2)
	fmt.Println(result)
}
