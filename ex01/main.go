package main
import (
		"fmt"
		"piscine"
)
func main() {
		fmt.Println(piscine.MakeRange(5, 10))
		fmt.Println(piscine.MakeRange(10, 5) == nil)
		fmt.Println(piscine.MakeRange(-10, -5))
		fmt.Println(piscine.MakeRange(-5, 5))
		fmt.Println(piscine.MakeRange(0, 5))
		fmt.Println(piscine.MakeRange(10, 5))
		fmt.Println(piscine.MakeRange(3, 3) == nil)
		fmt.Println(piscine.MakeRange(0, 0))
}