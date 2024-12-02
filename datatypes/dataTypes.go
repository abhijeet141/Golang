package main

import "fmt"

func main() {
	var x int16 = 225
	fmt.Println(x - 3)
	var y int16 = 83
	fmt.Println(x, y, x+y)
	a := 20.45
	b := 30.26
	c := b - a
	fmt.Printf("%f", c)
	str1 := "abhijeet"
	str2 := "abhijeet"
	str3 := "abhi"
	res1 := str1 == str2
	res2 := str1 == str3
	fmt.Println(res1, res2)
	str := "abhijeet"
	fmt.Println(len(str))
	var isAvailable bool = true
	var number []int = []int{1, 2, 3, 4, 5}
	fmt.Println(isAvailable)
	fmt.Println(number)
	nums := []int{1, 2, 3, 4}
	nums = append(nums, 5)
	fmt.Println(nums)
}
