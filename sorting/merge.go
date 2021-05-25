package sorting

func MergeSort(numbers []int) []int {
	mergeSortHelper(numbers, 0, len(numbers)-1)
	return numbers
}

func mergeSortHelper(numbers []int, begin int, end int) {
	// Base case
	if begin >= end {
		return
	}
	middle := (end-begin)/2 + begin
	mergeSortHelper(numbers, begin, middle)
	mergeSortHelper(numbers, middle+1, end)
	merge(numbers, begin, middle, end)
}

func merge(numbers []int, begin int, middle int, end int) {

	tmp := make([]int, end-begin+1)
	left := begin
	right := middle + 1
	counter := 0

	for left <= middle || right <= end {
		var value int
		if left <= middle && right > end {
			value = numbers[left]
			left++
		} else if right <= end && left > middle {
			value = numbers[right]
			right++
		} else if numbers[left] < numbers[right] {
			value = numbers[left]
			left++
		} else {
			value = numbers[right]
			right++
		}
		tmp[counter] = value
		counter++
	}
	counter = 0
	for i := begin; i <= end; i++ {
		numbers[i] = tmp[counter]
		counter++
	}

}
