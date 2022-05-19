package main
import "piscine"
func main() {
		// a := piscine.SplitWhiteSpaces("Hello how are you?")
		// a := piscine.SplitWhiteSpaces(" Hello how are you?")
		a := piscine.SplitWhiteSpaces(" Hello\n\nhow\tare  you? ")
		// a := piscine.SplitWhiteSpaces("")
		piscine.PrintWordsTables(a)
}

//a = []string{"Hello", "how", "are", "you?"}