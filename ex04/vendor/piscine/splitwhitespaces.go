package piscine

func CheckRune(c rune)bool{
	if c == '\t' || c == '\n' || c == '\v' || c == '\f' || c == '\r' || c == ' '{
		return true
	}
	return false
}

func SplitWhiteSpaces(s string) []string {
	var strs []string
	var hit_count int

	for i, check_word := range []rune(s){
		if i == 0 && CheckRune(check_word){
			strs = append(strs,"")
			hit_count = i+1
		} 
		if i > 0 && CheckRune(check_word){
			strs = append(strs, string([]rune(s)[hit_count:i]))
			hit_count = i+1
		}
	}
		strs = append(strs,string([]rune(s)[hit_count:]))
	return strs
}