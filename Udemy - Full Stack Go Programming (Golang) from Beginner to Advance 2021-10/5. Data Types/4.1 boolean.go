package main

import "fmt"

func main() {

	// decalare Boolean data type

	str1 := "Go Programming"
	str2 := "GO PROGRAMMING"
	str3 := "Go Programming"

	result1 := str1 == str2

	result2 := str1 == str3

	fmt.Println("Result 1 : ", result1)

	fmt.Println("Result 2 : ", result2)

	fmt.Printf("Type of Result 1 : %T\n", result1)
	fmt.Printf("Type of Result 2 : %T\n", result2)
}
