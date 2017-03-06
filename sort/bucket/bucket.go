/*
# bucket sort 桶排序

桶排序的基本思路是将数据根据计算规则来分组，并将数据依次分配到对应分组中。
分配时可能出现某分组里有多个数据，那么再进行分组内排序。
然后得到了有序分组，最后把它们依次取出来放到数组中即实现了整体排序。

此处分类即是桶，计算规则是 index = value * (n-1) /k,
(n为数据个数即桶个数，k为使全部数据分布在k*[0,1]空间中的正整数，k可取数据最大值)。
桶本身等价于一个二维数组。

注意点：是在插入桶的时候还是在插入桶之后进行桶内排序呢？我觉得入桶后比较好，
因为入桶后可以从桶内整体出发进行桶内排序，而不用挨个比较

动画演示 https://www.cs.usfca.edu/~galles/visualization/BucketSort.html

# runtime:k为桶数
-Worst:O(n^2)
-Best:Omega(n + k)
-Average:Theta(n + k)

# 稳定性

 */
package bucket

// import "log"
/*
桶内排序
 */
func sortInBucket(bucket []int) {//此处实现插入排序方式，其实可以用任意其他排序方式
	length := len(bucket)
	if length == 1 {return}

	for i := 1; i < length; i++ {
		backup := bucket[i]
		j := i -1;
		//将选出的被排数比较后插入左边有序区
		for  j >= 0 && backup < bucket[j] {//注意j >= 0必须在前边，否则会数组越界
			bucket[j+1] = bucket[j]//移动有序数组
			j -- //反向移动下标
		}
		bucket[j + 1] = backup //插队插入移动后的空位
	}
}
/*
获取数组最大值
 */
func getMaxInArr(arr []int) int{
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max{ max = arr[i]}
	}
	return max
}
/*
桶排序
 */
func Sort(arr []int) []int {
	//桶数
	num := len(arr)
	//k（数组最大值）
	max := getMaxInArr(arr)
	//二维切片
	buckets := make([][]int, num)

	//分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num-1) /max//分配桶index = value * (n-1) /k
		
		buckets[index] = append(buckets[index], arr[i])
	}
	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0{
			sortInBucket(buckets[i])

			copy(arr[tmpPos:], buckets[i])
			
			tmpPos += bucketLen
		}
	}

	return arr
}
