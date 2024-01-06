package sorting

func QuickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    pivot := medianOfThree(arr, 0, len(arr)/2, len(arr)-1)

    left, right := 0, len(arr)-1
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