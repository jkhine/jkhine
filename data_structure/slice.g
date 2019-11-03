package main

import "fmt"

func calc(in []int) ([]int, []int) {
	tmp := make([]int, len(in))
	copy(tmp, in)
	in = append(in, 4)
	return in, tmp
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	new, old := calc(a)
	fmt.Println(new)
	fmt.Println(old)
}
