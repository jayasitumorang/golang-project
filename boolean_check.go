/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main
import "fmt"

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	s := []string{"Jay", "aer", "Chris", "opip"}
	fmt.Println(contains(s, "corden")) // true
	fmt.Println(contains(s, "aer")) // false
}