package mergesort

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	length := len(input)
	if length <= 1 {
		return input
	}

	firstHalf := MergeSort(input[:length/2])
	secondHalf := MergeSort(input[length/2:])
	var i, j, k int
	output := make([]int, length)
	for i < len(firstHalf) {
		for j < len(secondHalf) && secondHalf[j] <= firstHalf[i] {
			output[k] = secondHalf[j]
			k++
			j++
		}
		output[k] = firstHalf[i]
		k++
		i++
	}
	for j < len(secondHalf) {
		output[k] = secondHalf[j]
		k++
		j++
	}
	return output
}
