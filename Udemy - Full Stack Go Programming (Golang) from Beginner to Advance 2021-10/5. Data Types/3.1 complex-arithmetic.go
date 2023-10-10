package main

import "fmt"

func main() {

	var a = 3 + 5i
	var b = 2 + 3i

	var result1 = a + b

	var result2 = b - a

	var result3 = a * b

	var result4 = a / b

	fmt.Println(result1, result2, result3, result4)
}
