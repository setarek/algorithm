package sort

func BubbleSort(input []int64) []int64 {

	for i := 0; i < len(input); i++ {
		for j := 1; j < len(input); j++ {
			if input[j] < input[j-1] {
				input[j], input[j-1] = input[j-1], input[j]
			}
		}

	}
	return input
}
