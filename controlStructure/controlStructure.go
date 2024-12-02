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
	fmt.Println("Enter age: ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	if age > 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}
}
