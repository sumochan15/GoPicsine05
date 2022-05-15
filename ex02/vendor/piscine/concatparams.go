package piscine

// func CountSize(args []string)(size int){
// 	for range args{
// 		size++
// 	}
// 	return
// }

// func MakeRange(args []string, size int)(new_args []string){
// 	new_args = make([]string, (size*2 - 1))
// 	return 
// }

// func AppendRange(new_args []string, args []string)[]string{
// 	size := CountSize(new_args)
// 	var re_args []string
// 	for i, j := 0, 0; i < size; i++ {
// 		if i%2 == 0 {
// 			new_args[i] = args[j]
// 			re_args = new_args 
// 			j++
// 		}
// 		if i%2 == 1{
// 			new_args[i] = "\n"
// 			re_args = new_args
// 		}
// 	}
// 	return re_args
// }

// func ConcatParams(args []string) string {
// 	size := CountSize(args)
// 	new_args := MakeRange(args, size)
// 	re_args := AppendRange(new_args, args)
// 	var cat_test string = "" 

// 	for _, strs := range re_args{
// 			cat_test += strs
// 	}
// 	return cat_test
// }

// func PrintStr(strs string)string{
// 	var r []rune
// 	var r_strs string
// 	for range strs{
// 		r = strs
// 		r_strs += r
// 	}
// 	return r_strs
// }

func ConcatParams(test []string)string {
	var strs string

	for _, args := range test{
		strs += args
		strs += "\n"
	}
	return strs
}