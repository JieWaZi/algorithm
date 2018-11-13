package sort

func QuickSort(numbers []int, left, right int) {
	if left >= right {  //跳出递归
		return
	}
	pivotPos := partition(numbers, left, right)
	QuickSort(numbers, left, pivotPos-1)  //左边再次递归
	QuickSort(numbers, pivotPos+1, right) //右边再次递归
}

func partition(numbers []int, left, right int) int {
	baseValue := numbers[left]
	i := left
	j := right

	for i < j {
		for i < j && numbers[j] > baseValue { //先从右边开始 当小于numbers[j] 就停止
			j--
		}
		for i < j && numbers[i] <= baseValue { //右边停止后 ，左边开始， 当大于numbers[i]就停止
			i++
		}
		swap(i, j,numbers) //把大的交换到右边，把小的交换到左边
	}
	swap(i, left, numbers) //最后把baseValue到中间

	return i
}

func swap(i, j int, a []int) {
	a[i], a[j] = a[j], a[i]
}