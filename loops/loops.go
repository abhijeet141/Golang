package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("Golang")
	}
	//whileloop
	i := 0
	for i < 3 {
		i += 2
	}
	fmt.Println(i)
	//range
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu"}
	for i := range days {
		fmt.Println(days[i])
	}
	//index-based
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v \n", index, day)
	}
	//break
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}
	//goto
	var x int = 0
Label1:
	for x < 8 {
		if x == 5 {
			x = x + 1
			goto Label1
		}
		fmt.Println(x)
		x = x + 1
	}
	//continue
	var y int = 0
	for y < 8 {
		if y == 5 {
			y = y + 2
			continue
		}
		fmt.Println(y)
		y
	}
}
