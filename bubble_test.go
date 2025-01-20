package algorithms

import (
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		slice []int
		want  []int
	}{
		{
			slice: []int{8, 2, 4, 0},
			want:  []int{0, 2, 4, 8},
		},
		{
			slice: []int{11, 9, 7, 5, 4, 1},
			want:  []int{1, 4, 5, 7, 9, 11},
		},
		{
			slice: []int{7, -4, 3, 0, 12, 9},
			want:  []int{-4, 0, 3, 7, 9, 12},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			old := tt.slice
			bubbleSort(tt.slice)
			if !reflect.DeepEqual(old, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", old, tt.want)
			}
		})
	}
}
