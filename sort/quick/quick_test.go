package quick

import (
	"fmt"
	"Algorithms-Learning-With-Go/sort/utils"
	
	"testing"
)


func TestQuickSort(t *testing.T) {
	list := utils.GetArrayOfSize(10)
			
	// Sort(list)
	HoareSort(list, 0, len(list)-1)
	// LomutoSort(list, 0, len(list)-1)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}


func benchmarkQuickSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	
	for i := 0; i < b.N; i++ {
		// Sort(list)
		// Sort(list)LomutoSort
		// LomutoSort(list, 0, len(list)-1)
		HoareSort(list, 0, len(list)-1)
	}
}

func BenchmarkQuickSort100(b *testing.B) { benchmarkQuickSort(100, b) }
func BenchmarkQuickSort1000(b *testing.B)   { benchmarkQuickSort(1000, b) }
func BenchmarkQuickSort10000(b *testing.B)  { benchmarkQuickSort(10000, b) }
func BenchmarkQuickSort100000(b *testing.B) { benchmarkQuickSort(100000, b) }
