package main

import "fmt"

func main() {
	fmt.Println("2D Slice")
	twoD := make([][]int, 3) //3rows
	for i := 0; i < 3; i++ {
		//each slice will have sub-slice
		twoD[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
