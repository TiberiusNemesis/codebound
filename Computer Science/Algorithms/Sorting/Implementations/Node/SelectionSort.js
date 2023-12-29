function selectionSort(array) {
    for (let currentElement = 0; currentElement < array.length; currentElement++) {
        let minIndex = currentElement
        for (let j = currentElement + 1; j < array.length; j++) {
            if (array[j] < array[minIndex]) {
                minIndex = j
            }
        }

        if (minIndex !== currentElement) {
            [array[currentElement], array[minIndex]] = [array[minIndex], array[currentElement]];
        }
    }

    return array
}