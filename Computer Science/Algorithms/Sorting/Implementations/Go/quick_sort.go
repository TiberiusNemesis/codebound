package sorting

func QuickSort(arr []int) {
    quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, start, end int) {
    for start < end {
        pivotIndex := medianOfThree(arr, start, start+(end-start)/2, end)
		pivotNewIndex := partition(arr, start, end, pivotIndex)

       quickSortHelper(arr, start, pivotNewIndex - 1)
	   start = pivotNewIndex + 1
    }
}

func partition(arr []int, start, end, pivotIndex int) int {
    pivot := arr[pivotIndex]

    arr[pivotIndex], arr[end] = arr[end], arr[pivotIndex]
	storeIndex := start

    for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[i], arr[storeIndex] = arr[storeIndex], arr[i]
			storeIndex++
		}
	}

	arr[storeIndex], arr[end] = arr[end], arr[storeIndex]
    return storeIndex
}


func medianOfThree(arr []int, a, b, c int) int {
    if arr[a] < arr[b] {
        if arr[b] < arr[c] {
            return arr[b]
        } else if arr[a] < arr[c] {
            return arr[c]
        } else {
            return arr[a]
        }
    } else {
        if arr[a] < arr[c] {
            return arr[a]
        } else if arr[b] < arr[c] {
            return arr[c]
        } else {
            return arr[b]
        }
    }
}