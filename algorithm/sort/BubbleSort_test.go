package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestBubbleSort(t *testing.T) {
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
			args: args{array: []int{5, 3, 2, 4, 1}},
			wants: wants{sortedArray: []int{1, 2, 3, 4, 5}, count: 8},
		},
	}

	for _, tt := range tests {
		count := BubbleSort(tt.args.array)

		assert.Equal(t, tt.wants.count, count)
		assert.Equal(t, tt.wants.sortedArray, tt.args.array)
	}
}