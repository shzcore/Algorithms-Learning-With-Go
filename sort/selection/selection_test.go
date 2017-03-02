package selection

import (
	"fmt"
	"Algorithms-Learning-With-Go/sort/utils"
	
	"testing"
)


func TestSelectionSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
			
	Sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func benchmarkSelectionSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	
	for i := 0; i < b.N; i++ {
		Sort(list)
	}
}

func BenchmarkSelectionSort100(b *testing.B) { benchmarkSelectionSort(100, b) }
func BenchmarkSelectionSort1000(b *testing.B)   { benchmarkSelectionSort(1000, b) }
func BenchmarkSelectionSort10000(b *testing.B)  { benchmarkSelectionSort(10000, b) }
func BenchmarkSelectionSort100000(b *testing.B) { benchmarkSelectionSort(100000, b) }
