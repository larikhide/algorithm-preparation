package quickSort

func QuickSort(arr []int, start, end int) {
	if start < end {
		pi := partion(arr, start, end)

		// Recursively sort elements before partition and after partition
		QuickSort(arr, start, pi-1)
		QuickSort(arr, pi+1, end)
	}
}

func partion(arr []int, start, end int) int {
	pivot := arr[end]
	i := start - 1

	for j := start; j < end; j++ {
		if arr[j] < pivot {
			i++

			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}
