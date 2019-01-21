package main

import (
	"fmt"
	// _ used to "forget" the import
	// Since the compiler won't let you import something
	// and not use it
	_ "net/http"
)

func main() {
	// Single line comment

	/*
		Block comment
	*/

	fmt.Println("Hello world!")

	// Explicit declaration
	var x int

	// Assignment
	x = 42

	// Implicit declaration and assignment
	y := 43

	// Show types using format verbs
	fmt.Printf("Variable x is of type: %T\n", x)
	fmt.Printf("Variable y is of type: %T\n", y)

	// Array declaration
	ages := [4]int{18, 19, 33, 41}

	// Loop over ages and print them
	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}

	// slice of names
	var names []string

	// Add a few names to the slice
	names = append(names, "Hunter")
	names = append(names, "Madison")
	names = append(names, "Logan")

	// Use of the _ variable to "forget" about the value when
	// using a variable, since the first variable is the index
	for _, val := range names {
		fmt.Println(val)
	}

	// This is a creation of an empty map
	// Assigning to this will cause a runtime
	// error
	// var colors map[string]string

	// Utilizing themake function
	// initializes it and allocated the space for the map
	colors := make(map[string]string)

	// Map value assignment
	colors["red"] = "#FF0000"

	// example of a function returning
	// more than 1 value
	sum, err := sumsUp()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sum)

	sum, err = sumsUp(1, 2, 3, 4, 5)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sum)

}

// Func isFoo() takes an argument "a" and
// returns a boolean value
func isFoo(a string) bool {
	if a == "foo" {
		return true
	}

	// Having this return statement here
	// is "The Go way" of dealing with
	// something like this
	return false
}

// Variadic function, takes as many args
// as you want and combines them into a slice
// No optional parameters :p
// but structs are a common way to get around this
func sumsUp(nums ...int) (int, error) {
	total := 0

	if len(nums) == 0 {
		return 0, fmt.Errorf("Nums must be longer than 0!")
	}

	for _, val := range nums {
		total += val
	}

	return total, nil
}
