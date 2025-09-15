package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println("Hello, Go Systems Programming!")
	// basic()
	// fmt.Println("3 + 4 =", add(3, 4))
	// result, err := divide(10, 2)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("10 / 2 =", result)
	// result, err = divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("10 / 0 =", result)
	// }

	// f := square()
	// fmt.Println(f()) // 1
	// fmt.Println(f()) // 4
	// fmt.Println(f()) // 9

	// loops()
	// slicesAndArrays()
	// maps()
	// structs()
	// sharks(

	// p := Point{1, 2}
	// q := Point{4, 6}
	// fmt.Printf("Distance between p and q is: %v\n", p.Distance(q))
	// fmt.Printf("Distance between p and q is: %v\n", Distance(p, q))

	goroutines()
	// Add a sleep to allow goroutines to finish
	var input string
	fmt.Scanln(&input)
	fmt.Println("Exiting main function")
}

// Package level variable
var msg string = "Variable type is inferred from the value on the right side"

// Package level function
func basic() {
	var x int

	var y int = 42

	z := 43 // short variable declaration

	x = 1
	y = 2
	z = x + 2*y + 3

	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)

	fmt.Println(msg)
}

func add(x, y int) int {
	return x + y
}

func divide(x, y int) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("division by zero")
	}

	return float64(x) / float64(y), nil
}

func square() func() int {
	var x int

	// Anonymous function
	// Definded at its point of use
	// Declare without a name
	return func() int {
		x++
		return x * x
	}
}

func loops() {
	// For loop
	for i := 0; i < 10; i++ {
		fmt.Print(".")
	}
	// While loop
	sum := 1
	for sum < 1000 {
		sum *= 2
	}
	fmt.Printf("The sum is %v\n", sum)
}

// Array: Fixed size, homogeneous (same type) collection of elements
// Slice: Dynamic size, homogeneous (same type) collection of elements
func slicesAndArrays() {
	var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array)
	fmt.Println(array[2:5]) // 3, 4, 5 [type slice]
	fmt.Println(array[5:])  // 6, 7, 8
	fmt.Println(array[:3])  // 1, 2, 3

	slice := make([]string, 3)
	slice[0] = "tic"
	slice[1] = "tac"
	slice[2] = "toe"
	fmt.Println(slice)
	slice = append(slice, "tom")
	slice = append(slice, "radar")
	fmt.Println(slice)
	for index, value := range slice {
		fmt.Printf("%v: %v\n", index, value)
	}
	fmt.Printf("Slice length = %v\n", len(slice))
}

// Map: Dynamic size, heterogeneous (key-value pairs) collection of elements
func maps() {
	// Declare a map whose keys have type string, and values have type int
	myMap := make(map[string]int)
	myMap["yellow"] = 1
	myMap["magic"] = 2
	myMap["amsterdam"] = 3
	fmt.Println(myMap)
	myMap["magic"] = 100
	delete(myMap, "amsterdam")
	fmt.Println(myMap)
	fmt.Printf("Map size = %v\n",
		len(myMap))
}

// Struct: Collection of fields
type Shark struct {
	Name string
	Age  int
}

// OOP paradigm
/*
	Object: a value or variable that contains both data and methods
	Method: a function that is associated with a particular type
		Method Declaration
		Similar to function declaration, but
		add an extra parameter between
		func and name. This will attach the
		function to the type of the
*/
// Go to file Point.go to see the struct Point and its methods

// go routines
func goroutines() {
	for i := 0; i < 10; i++ {
		// Print the number asynchronously
		go fmt.Printf("Printing %v in a goroutine\n", i)
	}
	// At this point the numbers may not have been printed yet
	fmt.Println("Launched the goroutines")
}
