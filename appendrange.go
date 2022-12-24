package piscine

func AppendRange(min, max int) []int {
	var result []int
	if max <= min {
		result = []int(nil)
	}
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}
