package sorting

import (
    "reflect"
    "testing"
)

func TestMergeSort(t *testing.T) {
    tests := []struct {
        name string
        arr  []int
        want []int
    }{
        {
            name: "Empty slice",
            arr:  []int{},
            want: []int{},
        },
        {
            name: "Single element",
            arr:  []int{1},
            want: []int{1},
        },
        {
            name: "Two elements",
            arr:  []int{2, 1},
            want: []int{1, 2},
        },
        {
            name: "Multiple elements",
            arr:  []int{4, 2, 3, 1},
            want: []int{1, 2, 3, 4},
        },
        {
            name: "Negative numbers",
            arr:  []int{-3, -1, -2},
            want: []int{-3, -2, -1},
        },
        {
            name: "Duplicates",
            arr:  []int{5, 2, 2, 1},
            want: []int{1, 2, 2, 5},
        },
        {
            name: "Already sorted",
            arr:  []int{1, 2, 3, 4, 5},
            want: []int{1, 2, 3, 4, 5},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := mergeSort(tt.arr)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("mergeSort(%v) = %v, want %v", tt.arr, got, tt.want)
            }
        })
    }
}

func TestMerge(t *testing.T) {
	tests := []struct {
        name string
        leftArr  []int
		rightArr []int
        want []int
    }{
        {
            name: "Empty slices",
            leftArr:  []int{},
			rightArr:  []int{},
            want: []int{},
        },
        {
            name: "Single element",
            leftArr:  []int{1},
			rightArr:  []int{2},
            want: []int{1, 2},
        },
    }

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := merge(tt.leftArr, tt.rightArr)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("merge(%v and %v) failed, expected %v, got %v", tt.leftArr, tt.rightArr, got, tt.want)
            }
        })
    }
}
