package sort

func merge(left, right []int64) []int64 {

	result := make([]int64, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result

}

func MergeSort(input []int64) []int64 {

	if len(input) <= 1 {
		return input
	}

	mid := len(input) / 2

	leftSide := input[0:mid]
	rightSide := input[mid:]

	leftSide = MergeSort(leftSide)
	rightSide = MergeSort(rightSide)

	return merge(leftSide, rightSide)
}
