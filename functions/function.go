package main

import "fmt"

//pass by value

func multiply(num1 int, num2 int) int {
	num1 *= 2
	return num1 * num2
}

// pass by refrence

func multiply2(a, b *int) int {
	*a = *a * 2
	return *a * *b
}

func sum(nums ...int) int {
	sumRes := 0
	for _, value := range nums {
		sumRes += value
	}
	return sumRes
}
func processedData(callback func(data string) string) {
	data := "Go is awesome"
	result := callback(data)
	fmt.Println(result)
}
func myFunc(a, b int) (int, int, int) {
	return a - b, a * b, a + b
}
func calculator(a, b int) (multiply int, divide int) {
	multiply = a * b
	divide = a / b
	return
}
func main() {
	x := 5
	y := 10
	result1 := multiply(x, y)
	fmt.Println(x)
	fmt.Println(result1)
	x1 := 5
	y1 := 10
	result2 := multiply2(&x1, &y1)
	fmt.Println(x1)
	fmt.Println(result2)
	//variadic function
	result3 := sum(1, 2, 3, 4, 5)
	fmt.Println(result3)
	result4 := sum(2, 6)
	fmt.Println(result4)
	//Anonymous function
	value := func() {
		fmt.Println("Welcome")
	}
	value()
	//somehow similar to callback function
	value2 := func(data string) string {
		return "Processed: " + data
	}
	value3 := func(data string) string {
		return "Modified: " + data
	}
	processedData(value2)
	processedData(value3)
	var myvar1, myvar2, myvar3 = myFunc(2, 4)
	fmt.Println(myvar1, myvar2, myvar3)
	//named-return value
	m, d := calculator(105, 7)
	fmt.Println(m, d)
}
