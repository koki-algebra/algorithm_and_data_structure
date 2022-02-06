package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrumpBubbleSort(t *testing.T) {
	type args struct {
		Trumps []Trump
	}
	
	tests := []struct {
		name string
		args args
		want []Trump
	}{
		{
			name: "1",
			args: args{
				Trumps: []Trump{
					{Suit: H, Value: 4},
					{Suit: C, Value: 9},
					{Suit: S, Value: 4},
					{Suit: D, Value: 2},
					{Suit: C, Value: 3},
				},
			},
			want: []Trump{
				{Suit: D, Value: 2},
				{Suit: C, Value: 3},
				{Suit: H, Value: 4},
				{Suit: S, Value: 4},
				{Suit: C, Value: 9},
			},
		},
	}

	for _, tt := range tests {
		// BubbleSort by number
		TrumpBubbleSort(tt.args.Trumps)

		assert.Equal(t, tt.want, tt.args.Trumps)
	}
}

func TestTrumpSelectionSort(t *testing.T) {
	type args struct {
		Trumps []Trump
	}
	
	tests := []struct {
		name string
		args args
		want []Trump
	}{
		{
			name: "1",
			args: args{
				Trumps: []Trump{
					{Suit: H, Value: 4},
					{Suit: C, Value: 9},
					{Suit: S, Value: 4},
					{Suit: D, Value: 2},
					{Suit: C, Value: 3},
				},
			},
			want: []Trump{
				{Suit: D, Value: 2},
				{Suit: C, Value: 3},
				{Suit: S, Value: 4},
				{Suit: H, Value: 4},
				{Suit: C, Value: 9},
			},
		},
	}

	for _, tt := range tests {
		// BubbleSort by number
		TrumpSelectionSort(tt.args.Trumps)

		assert.Equal(t, tt.want, tt.args.Trumps)
	}
}