package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert := assert.New(t)

	values := []int{}
	Quicksort(values)
	assert.ElementsMatch(values, []int{})

	values = []int{1}
	Quicksort(values)
	assert.ElementsMatch(values, []int{1})

	values = []int{1, 4}
	Quicksort(values)
	assert.ElementsMatch(values, []int{1, 4})

	values = []int{4, 4}
	Quicksort(values)
	assert.ElementsMatch(values, []int{4, 4})

	values = []int{1, 4, 3}
	Quicksort(values)
	assert.ElementsMatch(values, []int{1, 3, 4})

	values = []int{1, 4, 5, 3, 3}
	Quicksort(values)
	assert.ElementsMatch(values, []int{1, 3, 3, 4, 5})

}
