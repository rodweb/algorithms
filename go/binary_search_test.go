package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var list = []int{2, 3, 5, 7, 8}
	var tests = []struct {
		list []int
		item int
		want int
	}{
		{list, 1, -1},
		{list, 5, 2},
		{list, 7, 3},
		{list, 2, 0},
		{list, 6, -1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v,%d", tt.list, tt.item), func(t *testing.T) {
			got, _ := BinarySearch(tt.list, tt.item)
			if got != tt.want {
				t.Fatalf("Got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	_, err := BinarySearch([]int{1, 2, 3}, 4)
	if err == nil {
		t.Fatal("Got nil, want error")
	}
}
