package main

import (
	"fmt"
)

func sum1(mych chan int) {
	sumNum := 0
	for i := 1; i <= 20; i++ {
		sumNum += i
	}
	mych <- sumNum
}
func sum2(mych chan int, mych2 chan int) {
	sumNum := <-mych
	for i := 21; i <= 40; i++ {
		sumNum += i
	}
	mych2 <- sumNum
}
func sum3(mych2 chan int, mych3 chan int) {
	sumNum := <-mych2
	for i := 41; i <= 60; i++ {
		sumNum += i
	}
	mych3 <- sumNum
}
func sum4(mych3 chan int, mych4 chan int) {
	sumNum := <-mych3
	for i := 61; i <= 80; i++ {
		sumNum += i
	}
	mych4 <- sumNum
}
func sum5(mych4 chan int, mych5 chan int) {
	sumNum := <-mych4
	for i := 81; i <= 100; i++ {
		sumNum += i
	}
	mych5 <- sumNum
}
func main() {
	fmt.Println("Sum of numbers")
	mych := make(chan int)
	mych2 := make(chan int)
	mych3 := make(chan int)
	mych4 := make(chan int)
	mych5 := make(chan int)
	go sum1(mych)
	go sum2(mych, mych2)
	go sum3(mych2, mych3)
	go sum4(mych3, mych4)
	go sum5(mych4, mych5)

	fmt.Println(<-mych5)
}
