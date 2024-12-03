package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func calculateAverage(arr *[6]int, size int) int {
	var sum int
	for _, value := range arr {
		sum += value
	}
	return sum / size
}
func printSlice(slice []int) {
	fmt.Println(slice)
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for _, value := range arr {
		fmt.Println(value)
	}
	//2d array
	arr2 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(arr2[i][j])
		}
	}
	myarray := []string{"GFG", "gfg", "geeks",
		"GeeksforGeeks", "GEEK"} //slice dynamic in nature
	fmt.Println("Length of the array is:", len(myarray))
	myarray2 := [...]string{"GFG", "gfg", "geeks",
		"GeeksforGeeks", "GEEK"} //array fixed
	fmt.Println("Length of the array is:", len(myarray2))
	var destination [5]string = myarray2
	//in slice we can't directly assign
	fmt.Println(destination)
	scores := [6]int{1, 2, 3, 4, 5, 6}
	average := calculateAverage(&scores, len(scores))
	fmt.Println(average)
	//array and slice
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]
	fmt.Println(array)
	fmt.Println(slice)
	slice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice1)
	// var my_slice = make([]int, 4, 6)
	//slice is refrenced type so if we refer through a slice and make chage then array will also change
	s1 := [][]int{{12, 34},
		{56, 47},
		{29, 40},
		{46, 78},
	}
	fmt.Println(s1)
	slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"}
	slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}
	sort.Strings(slc1)
	sort.Ints(slc2)
	source := []int{1, 2, 3, 4, 5}
	destination1 := make([]int, len(source))
	count := copy(destination1, source)
	fmt.Println(source)
	fmt.Println(destination1)
	fmt.Println(count)
	a := []int{1, 2, 3, 4, 5}
	printSlice(a)
	a = append(a, 4)
	fmt.Println(a)
	intSlice := []int{42, 23, 16, 15, 8, 4}
	sort.Ints(intSlice)
	sort.IntsAreSorted(intSlice)
	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i] > intSlice[j]
	})
	data := []byte("a,b,c")
	parts := bytes.Split(data, []byte(","))
	fmt.Println(parts)
	str := "apple,banana,cherry"
	spliting := strings.Split(str, ",")
	fmt.Println(spliting)
}
