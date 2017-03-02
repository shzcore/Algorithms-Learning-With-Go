package module

import (
	"fmt"
	"Algorithms-Learning-With-Go/dummy/module"
	"testing"
)

/*
	单元测试
 */
func TestModule(t *testing.T) {
			
	Module()

	if false {
		fmt.Println("module")
		t.Error()
	}
}

func benchmarkModule(n int, b *testing.B) {
	Module()
}

/*
	benchmark 测试
 */
func BenchmarkQuickSort100(b *testing.B) { benchmarkModule(100, b) }
func BenchmarkQuickSort1000(b *testing.B) { benchmarkModule(1000, b) }
func BenchmarkQuickSort10000(b *testing.B) { benchmarkModule(10000, b) }
func BenchmarkQuickSort100000(b *testing.B) { benchmarkModule(100000, b) }