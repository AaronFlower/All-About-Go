package quicksort

// Quicksort use partition method to sort an array.
func Quicksort(data []int) {
	if len(data) > 1 {
		i := partition(data)
		Quicksort(data[:i])
		Quicksort(data[i+1:])
	}
}

func partition(data []int) int {
	i, j := -1, 0
	L := len(data)
	pivot := data[L-1]
	for ; j < L-1; j++ {
		if data[j] < pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	i++
	data[i], data[j] = data[j], data[i]
	return i
}
