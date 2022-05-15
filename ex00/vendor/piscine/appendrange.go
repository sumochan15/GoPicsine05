package piscine

func AppendRange(min, max int) []int {

	if max - min <= 0 {
		return nil
	}
	
	var rangenum []int

	for i := min; i < max; i++ {
		if i == min {
			rangenum = []int{min}
		} 
		if i != min {
		rangenum = append(rangenum,i)
		}
	}
	return rangenum	
}
