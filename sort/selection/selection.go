/*
# selection Sort 选择排序

类似冒泡排序，通过两层嵌套的循环来遍历、比较数组中元素。
不同的是选择排序在内层循环中不交换，而是选出最小值的位置，然后在外层循环中交换最小值到头部。
最终实现数组依次按照最小值排序。

选择排序即选择最值位置排序，效率差

# Runtime:

- Average: O(N^2)
- Worst: O(N^2)
*/
package selection

func Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
