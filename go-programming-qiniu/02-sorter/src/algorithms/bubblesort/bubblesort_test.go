package bubblesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	assert := assert.New(t)

	values := []int{}
	BubbleSort(values)
	assert.ElementsMatch(values, []int{})

	values = []int{1}
	BubbleSort(values)
	assert.ElementsMatch(values, []int{1})

	values = []int{1, 4}
	BubbleSort(values)
	assert.ElementsMatch(values, []int{1, 4})

	values = []int{4, 4}
	BubbleSort(values)
	assert.ElementsMatch(values, []int{4, 4})

	values = []int{1, 4, 3}
	BubbleSort(values)
	assert.ElementsMatch(values, []int{1, 3, 4})

	values = []int{1, 4, 5, 3, 3}
	BubbleSort(values)
	assert.ElementsMatch(values, []int{1, 3, 3, 4, 5})

}
