/*
# merge sort 归并排序

思路就是先分治后合并,这是种非原地（not-in-place）方式。

#runtime
-Worst:O(n log n)
-Best:O(n log n)
-Average:O(n log n)
#stability

 */
package merge

/*
合并
 */
func merge(leftPile, rightPile []int) []int {
  // 1
  leftIndex := 0
  rightIndex := 0

  // 2 
  orderedPile := []int{}

  // 3
  for leftIndex < len(leftPile) && rightIndex < len(rightPile) {
    if leftPile[leftIndex] < rightPile[rightIndex] {
      orderedPile = append(orderedPile,leftPile[leftIndex])
      leftIndex += 1
    } else if leftPile[leftIndex] > rightPile[rightIndex] {
      orderedPile = append(orderedPile,rightPile[rightIndex])
      rightIndex += 1
    } else {
      orderedPile = append(orderedPile,leftPile[leftIndex])
      leftIndex += 1
      orderedPile = append(orderedPile,rightPile[rightIndex])
      rightIndex += 1
    }
  }

  // 4
  for leftIndex < len(leftPile) {
    orderedPile = append(orderedPile,leftPile[leftIndex])
    leftIndex += 1
  }

  for rightIndex < len(rightPile) {
    orderedPile = append(orderedPile,rightPile[rightIndex])
    rightIndex += 1
  }
  return orderedPile
}
func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	//分
	leftPile := Sort(arr[:mid])
	rightPile := Sort(arr[mid:])
	//合
	return merge(leftPile, rightPile)
}