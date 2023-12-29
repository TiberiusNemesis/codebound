function bubbleSort (array) {
    let isSorted = true
    for (let index = 0; index < array.length; index++) {
        const element = array[index];
        const nextElement = array[index + 1]
        if (nextElement) {
            if (element > nextElement) {
                array[index] = nextElement
                array[index + 1] = element
                isSorted = false
            }
        }
    }

    if (!isSorted) {
        return bubbleSort(array)
    }

    return array
}

function efficientBubbleSort(array) {
    let n = array.length;
    let swapped;
    do {
        swapped = false;
        for (let i = 1; i < n; i++) {
            if (array[i - 1] > array[i]) {
                let temp = array[i - 1];
                array[i - 1] = array[i];
                array[i] = temp;
                swapped = true;
            }
        }
        n--; 
    } while (swapped);
    return array;
}