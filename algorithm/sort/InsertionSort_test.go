package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestInsertionSort(t *testing.T) {
	type args struct {
		A []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{[]int{5, 2, 4, 6, 1, 3}}, want: []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, InsertionSort(tt.args.A))
	}
}