# Pointers in Go

In Go, a pointer is a variable that stores the memory address of another variable. It points to the memory location where the value is stored.

## Why use pointers?

1. **Efficient memory management:**
    - It passes the memory address instead of copying the entire structure, which is more efficient.

2. **Modifying variables:**
    - When you pass a variable to a function via a pointer, you can modify the original value as the function operates on the actual memory address of the variable.

## How to use pointers

1. **Creating a pointer to a variable:**
    - Syntax: `var (variable name of the pointer) *(type of variable) = &(variable to which you want to point)`
    - Example:
      ```go
      var number int = 99
      var pointer *int = &number
      ```

2. **Accessing the value of the variable via pointer:**
    - Syntax: `*(variable name of the pointer)`
    - Example:
      ```go
      fmt.Println(*pointer) // Outputs: 99
      ```

3. **Modifying the value via pointer:**
    - Example:
      ```go
      *pointer = 100
      fmt.Println(number) // Outputs: 100
      ```