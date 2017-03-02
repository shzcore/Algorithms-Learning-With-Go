/*
# shell sort 希尔排序

#核心——通过设置下标增量来将数据分组，然后用插入排序法处理每个分组，再将下标增量减半。
接下来循环调用此过程直到下标增量为0。

又叫缩小增量排序算法，以科学家 DL．Shell 名字命名。

简直美爆了！小步长增量来插入排序时可以利用大步长增量的有序性，
利用了插入排序适合小数据量和基本有序的特点，充分发挥插入排序的优点，
使得整体用来处理大数据量的排序非常不错——虽然比快排还差点。


# runtime
- Worst:O(n^2)  O(n^1.5)  O(n((log^2)n))跟增量大小有关
- Best:O(n)
- Average:O(n^(1~2))

# stability
不稳定
 */
package shell

/*
	排序起始点与步长为变量而非常量0和1的时候乃是更为一般化的实现。
 */
func insertion(arr []int, start, gap int) {
	length := len(arr)

	for traverseVal :=  start + gap; traverseVal < length; traverseVal += gap {
		backup := arr[traverseVal]
		trackVal := traverseVal - gap
		for trackVal >= 0 && backup < arr[trackVal] {
			arr[trackVal + gap] = arr[trackVal]
			trackVal -= gap
		}
		arr[trackVal + gap] = backup
	}
}
func Sort(arr []int) []int {
	//设定步长增量
	gap := len(arr)/2
	//结束条件
	for gap > 0 {
		for pos := 0; pos<gap; pos ++{
			insertion(arr, pos, gap)
		}
		gap /= 2
	}

	return arr
}