// slices

package main

import "fmt"

func main() {
	var x_slice_var []int
	var x_slice_var_init = []int {1,2,3,4,5}
	x_slice_make := make([]int, 3)

	x_slice_var = append(x_slice_var, 1)
	x_slice_var = append(x_slice_var, 2)
	fmt.Println(x_slice_var)
	fmt.Println(x_slice_var_init)
	fmt.Println(x_slice_make)

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, i := range pow {
		fmt.Println(i)
	}
}
