/*
# Bubble Sort

Bubble sort is a sorting algorithm that is implemented by starting in the beginning of the array and swapping the first two elements only if the first element is greater than the second element. This comparison is then moved onto the next pair and so on and so forth. This is done until the array is completely sorted. The smaller items slowly “bubble” up to the beginning of the array.

# Runtime:
- Average: O(N^2)
- Worst: O(N^2)

# Memory:
- O(1)
*/
package bubble

// 冒泡排序
func Sort(arr []int) {

	for first := len(arr) - 1; ; first-- {
		//终止条件，也可通过外层循环来控制
		continued := false

		for second := first - 1; second >= 0; second-- {
			continued = true
			if arr[first] < arr[second] {
				arr[second], arr[first] = arr[first], arr[second]
			}
		}
		if continued == false {
			break
		}
	}

}
