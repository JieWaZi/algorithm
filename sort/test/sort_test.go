package test

import (
	"github.com/algorithm/sort"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkBubbleSort(b *testing.B) {
	numbers:=getRandomNumbers()
	sort.BubbleSort(numbers)
	if testSort(numbers) == false {
		b.Log("sort failed")
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	numbers:=getRandomNumbers()
	sort.InsertionSort(numbers)
	if testSort(numbers) == false {
		b.Log("sort failed")
	}
}

func BenchmarkChooseSort(b *testing.B) {
	numbers:=getRandomNumbers()
	sort.ChooseSort(numbers)
	if testSort(numbers) == false {
		b.Log("sort failed")
	}
}

func BenchmarkQuickSort(b *testing.B) {
	numbers:=getRandomNumbers()
	sort.QuickSort(numbers,0,len(numbers)-1)
	if testSort(numbers) == false {
		b.Log("sort failed")
	}
}

func BenchmarkMergeSort(b *testing.B) {
	numbers:=getRandomNumbers()
	numbers=sort.MergeSort(numbers)
	if testSort(numbers) == false {
		b.Log("sort failed")
	}
}


func getRandomNumbers() []int {
	numbers := make([]int, 0, 1000000)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000000; i++ {
		numbers = append(numbers, rand.Intn(100000))
	}

	return numbers
}

func testSort(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}