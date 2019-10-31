package main

import "fmt"

func add(args ...int) int {
	var total int
	for _, i := range args {
		total += i
	}
	return total
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}


func main() {
	fmt.Println(add(1,2,3,4,5))

	inside_add := func(args ...int) int {
		var tot int
		for _, x := range args {
			tot += x
		}
		return tot
	}

	fmt.Println(inside_add(1,2,3,4))

	defer second()
	first()
}
