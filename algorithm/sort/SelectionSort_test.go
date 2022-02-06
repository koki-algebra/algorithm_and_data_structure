package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		array []int
	}

	type wants struct {
		sortedArray []int
		count int
	}

	tests := []struct {
		name string
		args args
		wants wants
	}{
		{
			name: "1",
			args: args{array: []int{5, 6, 4, 2, 1, 3}},
			wants: wants{sortedArray: []int{1, 2, 3, 4, 5, 6}, count: 4},
		},
	}

	for _, tt := range tests {
		count := SelectionSort(tt.args.array)

		assert.Equal(t, tt.wants.count, count)
		assert.Equal(t, tt.wants.sortedArray, tt.args.array)
	}
}