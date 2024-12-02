package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	diceNumber := r.Intn(6) + 1
	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 step")
	case 2:
		fmt.Println("You can move 2 step")
	case 3:
		fmt.Println("You can move 3 step")
	case 4:
		fmt.Println("You can move 4 step")
	case 5:
		fmt.Println("You can move 5 step")
		fallthrough
	case 6:
		fmt.Println("You can move 6 step")
	default:
		fmt.Println("Wrong Move!!")
	}
}
