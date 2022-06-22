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
