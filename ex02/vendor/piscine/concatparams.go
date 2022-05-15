package piscine

func ConcatParams(test []string)string {
	var strs string

	for _, args := range test{
		strs += args
		strs += "\n"
	}
	return strs
}