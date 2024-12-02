package main

import "fmt"

func main() {
	var r rune = 'a'
	fmt.Println(r)
	fmt.Printf("%c", r)
	r1 := 'a'
	s := string(r1)
	fmt.Println(s)
	r2 := 'a'
	r2 = 'b'
	fmt.Println(r2)
}
