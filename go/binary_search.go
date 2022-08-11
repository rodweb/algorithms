package binary_search

import (
	"errors"
	"fmt"
)

func BinarySearch(list []int, item int) (int, error) {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return mid, nil
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
		fmt.Printf("item=%d,mid=%d,low=%d,high=%d\n", item, mid, low, high)
	}
	return -1, errors.New("Item not found.")
}
