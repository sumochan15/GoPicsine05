package main
import (
		"fmt"
		"piscine"
)
func main() {
fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello\thow\nare\fyou?"))
fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how  "))
fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" Hello how are you? "))
fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(""))
fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" "))
fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("     Hello   how are you?"))
}

