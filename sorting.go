package main

import (
	"fmt"
	"sort"
)

func main() {
	
	scl1 := []int{400, 600, 100, 300, 500, 200, 900}
	scl2 := []int{-23, 567, -34, 67, 0, 12, -5}
	

	z := append([]int{}, append(scl1, scl2...)...)
	
	sort.Ints (z)
	
	fmt.Println(z)
}
