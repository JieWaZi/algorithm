package sort

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j],numbers[j+1] = numbers[j+1],numbers[j]
			}
		}
	}
}