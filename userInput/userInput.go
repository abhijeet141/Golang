package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println(input)
	fmt.Printf("Type of input is %T", input)
}
