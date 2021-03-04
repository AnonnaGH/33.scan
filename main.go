package main

import "fmt"

func main() {

	var a, b, c float32

	fmt.Println(" Enter 1st number:")
	fmt.Scan(&a)

	fmt.Println(" Enter 2nd number:")
	fmt.Scan(&b)

	c = a / b

	fmt.Printf("%.2f",c)

}
