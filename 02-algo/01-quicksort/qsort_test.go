package qsort

import "testing"

func TestQuicksort0(t *testing.T) {
	data := []int{1}
	sdata := []int{1}
	sort(data)
	index := make([]int, len(data))
	for i := range index {
		if data[i] != sdata[i] {
			t.Errorf(" Expect data[%d](%d) = sdata[%d](%d)", i, data[i], i, sdata[i])
		}
	}
}

func TestQuicksort1(t *testing.T) {
	data := []int{1, 4, 5, 8, 3, 4}
	sdata := []int{1, 3, 4, 4, 5, 8}
	sort(data)
	index := make([]int, len(data))
	for i := range index {
		if data[i] != sdata[i] {
			t.Errorf(" Expect data[%d](%d) = sdata[%d](%d)", i, data[i], i, sdata[i])
		}
	}
}

func TestQuicksort2(t *testing.T) {
	data := []int{-1, 4, -5, -8, -8, 3, 4}
	sdata := []int{-8, -8, -5, -1, 3, 4, 4}
	sort(data)
	index := make([]int, len(data))
	for i := range index {
		if data[i] != sdata[i] {
			t.Errorf(" Expect data[%d](%d) = sdata[%d](%d)", i, data[i], i, sdata[i])
		}
	}
}
