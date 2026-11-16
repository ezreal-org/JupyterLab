package main

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "nil slice", nums: nil, want: nil},
		{name: "empty slice", nums: []int{}, want: []int{}},
		{name: "single element", nums: []int{42}, want: []int{42}},
		{name: "two elements", nums: []int{2, 1}, want: []int{1, 2}},
		{name: "already sorted", nums: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{name: "reverse order", nums: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
		{name: "unsorted", nums: []int{5, 3, 8, 1, 2}, want: []int{1, 2, 3, 5, 8}},
		{name: "duplicates", nums: []int{3, 1, 3, 2, 1}, want: []int{1, 1, 2, 3, 3}},
		{name: "all equal", nums: []int{7, 7, 7, 7}, want: []int{7, 7, 7, 7}},
		{name: "negative numbers", nums: []int{-3, -1, -5, -2}, want: []int{-5, -3, -2, -1}},
		{name: "mixed signs and zero", nums: []int{3, -1, 0, -2, 3}, want: []int{-2, -1, 0, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			BubbleSort(nums)
			if !slices.Equal(nums, tt.want) {
				t.Errorf("BubbleSort(%v) = %v, want %v", tt.nums, nums, tt.want)
			}
		})
	}
}
