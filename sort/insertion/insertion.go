/*
# insertion sort 插入排序

插入排序可以原地(in-place)也可以非原地(not-in-place)，核心步骤是在无序数组中选出被排数，
然后通过比较来移动有序数组并得到空位，最后将被排数插入空位。

插入排序非常适合少量或者已基本有序的数据，数据量变大时 O(n^2) 的复杂度会很恐怖（毕竟有两层循环）

非原地方式是申请新的数组然后把被排数插入新数组，原地方式是备份被排数然后从头部（或尾部）比较、移动来建立有序数组

# runtime
- Worst:O(n^2)
- Best:
- Average:O(n^2)

 */
package insertion

/*
	in-place 方式
 */
func Sort(arr []int) []int {
	length := len(arr)
	if length == 1 {return arr}

	for i := 1; i < length; i++ {
		backup := arr[i]
		j := i -1;
		//将选出的被排数比较后插入左边有序区
		for  j >= 0 && backup < arr[j] {//注意j >= 0必须在前边，否则会数组越界
			arr[j+1] = arr[j]//移动有序数组
			j -- //反向移动下标
		}
		arr[j + 1] = backup //插队插入移动后的空位
	}
	return arr
}