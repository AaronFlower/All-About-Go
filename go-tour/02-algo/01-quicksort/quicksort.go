package qsort

func sort(data []int) {
	num := len(data)
	if num <= 1 {
		return
	}
	i := partition(data)
	sort(data[:i])
	if i+1 < num {
		sort(data[i+1:])
	}
}

func partition(data []int) int {
	num := len(data)
	pivot := data[num-1]
	i := -1
	for j := 0; j < num; j++ {
		if data[j] < pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[i+1], data[num-1] = data[num-1], data[i+1]
	return i + 1
}
