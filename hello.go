package main

import (
	"fmt"
	"math"
	"math/rand"
)

func square() {
	fmt.Println("Square root of 45 is ", math.Sqrt(45))
}

func add(x,y float64) float64 {
	return x + y
}

func main() {
	var name = "Welcome to Go!\n"
	var age = 26
	fmt.Printf(name)
	fmt.Printf("My name is Chathura and my age is %d \n", age)
	fmt.Println("My name is Chathura and my age is", age)

	square()

	fmt.Printf("Random number %d \n", rand.Intn(100))

	var num1,num2 float64 = 30, 15
	fmt.Println(add(num1, num2))

	num5, num6 := 45, 65
	var num3,num4 float64 = float64(num5),float64(num6)
	fmt.Println(add(num3, num4))


	x := 10
	a := &x
	fmt.Println("Variable \"x\" memory address(&x) is",a)
	fmt.Println("Value of *a is",*a)
}
