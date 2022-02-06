package algorithm

import "testing"

type args struct {
	n int
	R []int
}

func TestMaximumProfit_test(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{n: 6, R: []int{5, 3, 1, 3, 4, 3}}, want: 3},
	}

	for _, tt := range tests {
		if got := MaximumProfit(tt.args.n, tt.args.R); got != tt.want {
			t.Errorf("error: got = %v, want = %v", got, tt.want)
		}
	}
}