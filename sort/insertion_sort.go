package sort

func InsertionSort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			j := i - 1
			temp := numbers[i]
			for j >= 0 && temp < numbers[j] {
				numbers[j+1] = numbers[j]
				j--
			}
			numbers[j+1] = temp
		}
	}
}
