package sort

func InsertionSort(input []int64) []int64 {
	for i := 1; i < len(input); i++ {
		j := i
		for j > 0 {
			if input[j] < input[j-1] {
				current := input[j]
				input[j] = input[j-1]
				input[j-1] = current
			}
			j--
		}
	}
	return input
}
