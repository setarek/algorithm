package sort

func partition(input []int64, low, high int64) ([]int64, int64) {
	pivot := input[high]
	i := low
	for j := low; j < high; j++ {
		if input[j] < pivot {
			input[i], input[j] = input[j], input[i]
			i++
		}
	}
	input[i], input[high] = input[high], input[i]
	return input, i
}

func QuickSort(input []int64, low, high int64) []int64 {
	if low < high {
		var p int64
		input, p = partition(input, low, high)
		input = QuickSort(input, low, p-1)
		input = QuickSort(input, p+1, high)
	}
	return input
}
