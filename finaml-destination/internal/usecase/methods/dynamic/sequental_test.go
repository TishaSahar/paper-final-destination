package dynamic

import "testing"

func TestMinCostAssignment(t *testing.T) {
	type args struct {
		tree [][]int
		MW   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "N = 2",
			args: args{
				tree: [][]int{{0, 1}, {1, 0}},
				MW:   []int{1, 0},
			},
			want: 0,
		},
		{
			name: "N = 3",
			args: args{
				tree: [][]int{{0, 1, 2}, {2, 1, 0}, {2, 0, 1}},
				MW:   []int{1, 0, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCostAssignment(tt.args.tree, tt.args.MW); got != tt.want {
				t.Errorf("MinCostAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}
