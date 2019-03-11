package bubblesort

// BubbleSort use bubble method to sort an array.
func BubbleSort(data []int) {
	l := len(data)
	minIndex := 0
	for i := 0; i < l; i++ {
		minIndex = i
		for j := i + 1; j < l; j++ {
			if data[minIndex] > data[j] {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
