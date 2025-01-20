package algorithms

import (
	"testing"
)

func Test_binSearch(t *testing.T) {
	tests := []struct {
		slice  []int
		target int
		want   int
	}{
		{
			slice:  []int{4, 5},
			target: 4,
			want:   0,
		},
		{
			slice:  []int{2, 4, 6, 8},
			target: 6,
			want:   2,
		},
		{
			slice:  []int{2, 4, 4, 4},
			target: 4,
			want:   1,
		},
		{
			slice:  []int{3, 5, 7, 9},
			target: 1,
			want:   -1,
		},
		{
			slice:  []int{11},
			target: 11,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := binSearch(tt.slice, tt.target); got != tt.want {
				t.Errorf("binSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
