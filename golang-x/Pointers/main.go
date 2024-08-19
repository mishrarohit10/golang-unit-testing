package main

import "fmt"

func main() {
	fmt.Println("Lets go pointers")

	// declaration
	var number int = 99

	// var (variable name of the pointer) *(type of variable) = &(variable to which you want to point)

	// creating a pointer
	var pointer *int = &number

	fmt.Println("Value of number -> ", number)
	fmt.Println("Pointer pointer (memory address of number) ->", pointer)

	// To access the value of the variable via pointer use *(variable name of the pointer)

	// to get value of pointer, dereference the pointer to get value stored at memory address
	fmt.Println("Value of the pointer ->", *pointer)

	// can modify the value via pointer as well
	*pointer = 34

	fmt.Println("Modified value via pointer -> ", pointer)
}
