package main

import (
	"fmt"
	"github.com/jkhine/hello_main/hello_func"
	"github.com/jkhine/stringutil"
)

func main() {
	fmt.Println(hello_func.Hello())
	fmt.Println(hello_func.Hello_cap())
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("Hello, world.")
}
