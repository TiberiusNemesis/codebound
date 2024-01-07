package sorting

import (
    "reflect"
    "testing"
)

func TestQuickSort(t *testing.T) {
    testCases := []struct {
        name     string
        input    []int
        expected []int
    }{
        {
            name:     "empty slice",
            input:    []int{},
            expected: []int{},
        },
        {
            name:     "single element slice",
            input:    []int{1},
            expected: []int{1},
        },
        {
            name:     "two elements slice",
            input:    []int{2, 1},
            expected: []int{1, 2},
        },
        {
            name:     "multiple elements slice",
            input:    []int{4, 1, 3, 2},
            expected: []int{1, 2, 3, 4},
        },
        {
            name:     "negative numbers slice",
            input:    []int{-3, -1, -2},
            expected: []int{-3, -2, -1},
        },
        {
            name:     "duplicate elements",
            input:    []int{5, 3, 3, 2},
            expected: []int{2, 3, 3, 5},
        },
        {
            name:     "already sorted slice",
            input:    []int{1, 2, 3, 4, 5},
            expected: []int{1, 2, 3, 4, 5},
        },
        {
            name:     "reverse sorted slice",
            input:    []int{5, 4, 3, 2, 1},
            expected: []int{1, 2, 3, 4, 5},
        },
        {
            name:     "large range of numbers",
            input:    []int{1412, -1678, 2596, 965, -64, 35, 0, 1},
            expected: []int{-1678, -64, 0, 1, 35, 965, 1412, 2596},
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            QuickSort(tc.input)
            if !reflect.DeepEqual(tc.input, tc.expected) {
                t.Errorf("QuickSort a(n) %s: got %v, want %v", tc.name, tc.input, tc.expected)
            }
        })
    }
}
