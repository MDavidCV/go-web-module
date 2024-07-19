package utility

func CountValues[T1 comparable](s []T1) map[T1]int {
	counts := make(map[T1]int)
	for _, value := range s {
		counts[value]++
	}
	return counts
}
