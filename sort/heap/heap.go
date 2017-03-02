/*
# heap Sort 堆排序

利用“堆”这种数据结构来排序，分为最大堆、最小堆排序。堆是种完全二叉树，特点是父子有序兄弟无序（可能有序可能乱序）。
排序过程即循环构建堆并依次排列堆的根节点的过程。

关键过程为建堆、堆顶与无序区尾部叶子交换

# Runtime:

- Average: O(n lg n)
- Worst: O(n lg n)
- Best: O(n lg n)
*/
package heap
// import "log"
/*
	golang 自带 container/heap，此处是对固定类型的一种简单实现最大堆
 */
func BuildHeap(arr []int, length int){
	if length == 1 {
		return
	}
	//最大非叶节点偏移量（序号为length / 2）
	maxBranch := length / 2 - 1
	//构造全部子树
	for i := maxBranch; i >= 0; i-- {
		//左孩子偏移量(序号为 2 * (length / 2))
		lChild := 2 * i + 1
		//右孩子偏移量(序号为 2 * (length / 2) + 1)
		rChild := lChild + 1
		//临时最大值偏移量
		tmpMax := i

		// if rChild > length { break }//不可以，左孩子要与父结点比较

		//将三个节点构造成堆（父结点最大）
		if arr[lChild] > arr[tmpMax] {//左孩子不需判断是否溢出
			tmpMax = lChild
		}
		if rChild <= length -1 && arr[rChild] > arr[tmpMax] {
			tmpMax = rChild
		}
		if tmpMax != i {
			arr[i], arr[tmpMax] = arr[tmpMax], arr[i]
		}
	}
}
func Sort(arr []int) []int {

	length := len(arr)
	for i := 0; i < length; i++ {
		//将需要重新建堆的长度减短
		lastMessLen := length-i

		BuildHeap(arr, lastMessLen)
		//将堆顶（根节点）与无序区的最后叶子节点交换
		if i < length{
			arr[0], arr[lastMessLen - 1] = arr[lastMessLen - 1],arr[0]
		}
	}	
	return arr
}