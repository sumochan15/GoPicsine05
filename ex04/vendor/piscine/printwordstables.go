package piscine

import "ft"

func PrintWordsTables(a []string) {
	for _, str := range a{
		PrintStr(str)
		ft.PrintRune('\n')
	}
}

func PrintStr(str string){
	for _, r := range str{
		ft.PrintRune(r)
	}
}