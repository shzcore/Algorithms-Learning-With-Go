/*
# merge sort 归并排序

思路就是先分治后合并,这是种非原地（not-in-place）方式。
实现过程是通过递归来从上到下分割，在最深层实现两数比较，
然后逐层上升连续将左右分区中最小值写入结果数组，最终实现整体排序。

#runtime
-Worst:O(n log n)
-Best:O(n log n)
-Average:O(n log n)

#stability
稳定
 */
package merge

/*
	合并
 */
func merge(leftPile, rightPile []int) []int {
  // 1 左右偏移量
  leftIndex := 0
  rightIndex := 0

  // 2 用以存放结果的数组
  sortedPile := []int{}

  // 3 递归最深层会是简单的两数比较，然后返回两数的有序数组；
  // 逐层上升来比较时则遍历左右两个有序数组，并选择其中最小的数据合并到结果数组中；
  // 遍历结束后直接将剩余有序值全部合并到结果数组，最终实现两个有序数组的整体有序
  for leftIndex < len(leftPile) && rightIndex < len(rightPile) {
    if leftPile[leftIndex] < rightPile[rightIndex] {
      sortedPile = append(sortedPile,leftPile[leftIndex])
      leftIndex += 1
    } else if leftPile[leftIndex] > rightPile[rightIndex] {
      sortedPile = append(sortedPile,rightPile[rightIndex])
      rightIndex += 1
    } else {
      sortedPile = append(sortedPile,leftPile[leftIndex])
      leftIndex += 1
      sortedPile = append(sortedPile,rightPile[rightIndex])
      rightIndex += 1
    }
  }

  // 4 如果左边剩下高值
  for leftIndex < len(leftPile) {
    sortedPile = append(sortedPile,leftPile[leftIndex])
    leftIndex += 1
  }
  // 5 如果右边剩下高值
  for rightIndex < len(rightPile) {
    sortedPile = append(sortedPile,rightPile[rightIndex])
    rightIndex += 1
  }
  return sortedPile
}
func Sort(arr []int) []int {
	//如果只剩一个元素则直接返回（递归方式最终都会产生最简数据单位，然后再逐层上升来运算）
	if len(arr) <= 1 { return arr }
	//否则对半分割，递归之后都会触发以上条件并返回，然后调用合并方法进行排序
	mid := len(arr) / 2
	//分
	leftPile := Sort(arr[:mid])
	rightPile := Sort(arr[mid:])
	//合
	return merge(leftPile, rightPile)
}