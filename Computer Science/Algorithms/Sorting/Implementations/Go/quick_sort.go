package sorting

func QuickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
	
    left, right := 0, len(arr)-1

    pivot := arr[len(arr)/2]

    for left <= right {
        for arr[left] < pivot {
            left++
        }
        for arr[right] > pivot {
            right--
        }
        if left <= right {
            arr[left], arr[right] = arr[right], arr[left]
            left++
            right--
        }
    }

    if 0 < right {
        QuickSort(arr[:right+1])
    }
    if left < len(arr) {
        QuickSort(arr[left:])
    }

    return arr
}