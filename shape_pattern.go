package main

import (
	"fmt"
	"strings"
)

func main() {
    var n int 
    fmt.Print("Enter the N= ")
	fmt.Scanln(&n)
	for i := n; i >= 1 && i <= 10; i-- {
		fmt.Printf("%s\n", strings.Repeat("#", i))
	}
}