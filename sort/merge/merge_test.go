package merge

import (
	"fmt"
	"Algorithms-Learning-With-Go/sort/utils"
	"testing"
)


func TestMergeSort(t *testing.T) {
	list := utils.GetArrayOfSize(10)
			
	list = Sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}


func benchmarkMergeSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	
	for i := 0; i < b.N; i++ {
		list = Sort(list)
	}
}

func BenchmarkMergeSort100(b *testing.B) { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)   { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B)  { benchmarkMergeSort(10000, b) }
func BenchmarkMergeSort100000(b *testing.B) { benchmarkMergeSort(100000, b) }
