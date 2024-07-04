function bubbleSort(arr) {
    let n = arr.length;
    let swapped;
    do {
        swapped = false;
        for (let i = 1; i < n; i++) {
            if (arr[i - 1] > arr[i]) {
                // Swap arr[i - 1] and arr[i]
                let temp = arr[i - 1];
                arr[i - 1] = arr[i];
                arr[i] = temp;
                swapped = true;
            }
        }
        // Reduce the range for the next iteration since the last element is already sorted
        n--;
    } while (swapped);
    return arr;
}

// Example usage
let array = [64, 34, 25, 12, 22, 11, 90];
console.log("Sorted array: ", bubbleSort(array));
