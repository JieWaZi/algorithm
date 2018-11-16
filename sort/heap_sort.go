package sort

/*
 **
 */
func HeapSort(values []int) {
	//构建一个大顶堆
	buildBigHeap(values)
	for i := len(values)-1; i > 0; i-- {
		values[0], values[i] = values[i], values[0]
		adjustHeap(values[:i], 0)
	}
}

func buildBigHeap(values []int) {
	for i := len(values)/2-1; i >= 0; i-- {
		adjustHeap(values, i)
	}
}


//调整堆，找出与index相比最大的指，并互换
func adjustHeap(values []int, index int) {
	node := index
	for node < len(values) {
		var child int = 0
		if 2*node+1 < len(values) {
			child = 2*node + 1
			if 2*node+2 < len(values) && values[2*node+2] > values[2*node+1] {
				child = 2*node + 2
			}
		}
		if child > 0 && values[child] > values[node] {
			values[node], values[child] = values[child], values[node]
			node = child
		} else {
			break
		}
	}
}
