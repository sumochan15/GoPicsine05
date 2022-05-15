package main
import (
		"fmt"
		"piscine"
)
func main() {
		test := []string{"Hello", "how", "are", "you?"}
		// test := []string{"","Hello", "how", "are", "you?"}
		// test := []string{"",""}
		fmt.Println(piscine.ConcatParams(test))
}