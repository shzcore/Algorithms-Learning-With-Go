/*
测试数据：
5 10
4 9
3 6
5 1
2 4
5 1
4 9
4 20
3 6
4 20
2 4
5 10
2 6
2 3
6 5
5 4
4 6
 */

package zeroonebag

import (
	"fmt"
	"testing"
)

func getTestArr() []map[string]int{

	return []map[string]int{
		map[string]int{"value":5,"weight":10},
		map[string]int{"value":4,"weight":9},
		map[string]int{"value":3,"weight":6},
		map[string]int{"value":5,"weight":1},
		map[string]int{"value":2,"weight":3},
		map[string]int{"value":5,"weight":1},
		map[string]int{"value":6,"weight":1},
		map[string]int{"value":4,"weight":9},
		map[string]int{"value":4,"weight":20},
		map[string]int{"value":2,"weight":4},
		map[string]int{"value":5,"weight":10},
		map[string]int{"value":2,"weight":6},
		map[string]int{"value":2,"weight":3},
		map[string]int{"value":6,"weight":2},
		map[string]int{"value":5,"weight":4},
		map[string]int{"value":4,"weight":6},
	}
}
func TestOneZeroBag(t *testing.T) {
	data := getTestArr()
	maxValue := run(data, 6)
	
	if maxValue != 22{
		fmt.Printf("得到最大值%d,应为%d \n",maxValue,22)
		t.Error()
	}
}