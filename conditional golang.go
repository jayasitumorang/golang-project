package main

import "fmt"

func main() {

	var eonum int

	fmt.Print("Enter the Number to find Even or Odd = ")
	fmt.Scanln(&eonum)
	if eonum%2 == 0 && eonum < 6 {
		fmt.Println("\nnot cool")
	} else if eonum%2 == 0 && eonum >= 6 {
		fmt.Println("\ncool")
	} else {
		fmt.Println("\nCool")
	}
}
