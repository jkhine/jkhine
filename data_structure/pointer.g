package main

import "fmt"

func ref(x *int) {
	fmt.Println("\nPass by ref")
	*x = 100
	fmt.Printf("value = %d, address = %+v\n\n", x, &x)
}

func val(x int) {
	fmt.Println("Pass by value")
	x = 100
	fmt.Printf("value = %d, address = %+v\n\n", x, &x)
}

func main() {
	var a int
	a = 5
	val(a)
	fmt.Println("after pass by value, number is ", a)
	ref(&a)
	fmt.Println("after pass by ref, number is ", a)
}
