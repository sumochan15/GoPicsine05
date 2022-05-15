package piscine

func countRange(min, max int)(n int){
	n = max - min
	return
}

func MakeRange(min, max int) []int {
	n := countRange(min, max)
	if n <= 0 {
		return nil
	}
	rangenum := make([]int,n)

	for i := 0; i < n; i++ {
		rangenum[i] = min
		min++
	}
	return rangenum
}