/*
# quick sort 快速排序

快速排序有几种实现方式，不过核心操作是选出一个基点比较值，通过一次遍历得出比它小的A和比它大的B无序数组，
然后通过递归方式分别操作A与B,并将低值A结果放到基点左边，高值B结果放到基点右边。

全过程都是为了基准点pivot归位。

# runtime
-Worst:O(n ^ 2)
-Best:O(n lg n)
-Average:O(n lg n)
*/
package quick
/*
	简单快排：
	递归模板返回小于、等于、大于基点比较元素的三个数组（低值数组、等值数组、高值数组），
	基点元素是数组中间值，递归完后将三个数组合并即是结果。
	
	评价——
	数据量过百后性能就很差。
	
	关键点——
	基点比较元素（pivot）
	低值数组
	等值数组
	高值数组
 */ 
func Sort(list []int) []int{

	if len(list) <= 1 {
		return list
	}

	//[[----递归模板区

	pivot := list[len(list)/2]
	
	partitionFunc :=  func (arr []int, pivot int) ([]int, []int, []int){//基点分割函数
		lessArr := make([]int, 0)
		greaterArr := make([]int, 0)
		equalArr := make([]int, 0)
		
		for _, value := range arr {
			switch {
				case value < pivot:
					lessArr = append(lessArr,value)
				case value > pivot:
					greaterArr = append(greaterArr,value)
				default:
					equalArr = append(equalArr,value)
			}
		}
		return lessArr, equalArr, greaterArr
	}
	less, equal, greater := partitionFunc(list, pivot)
	//--------------]]

	copy(list, Sort(less))
	copy(list[len(less):], equal)
	copy(list[(len(less)+len(equal)):], Sort(greater))

	return list
}

/*
	老毛桃 (Lomuto) 快排：
	以最右位置为基点比较元素，递归模板中不返回排序完的数组，而是返回遍历后基点比较元素应该所在的位置下标p，
	左边分区（包含小于等于基点元素的元素）是相对p值的低值数组，右边分区（包含大于基点元素的元素）
	是相对p值的高值数组。

	评价——
	比霍尔快排简单但是性能稍差。

	关键点——
	基点比较元素（pivot）
	遍历索引
	低值数组锚点
*/
func LomutoSort(list []int, low int, high int){
	if low >= high {return}

	//[[----递归模板区
	pivot := list[high]
	i := low
	for j := low+1; j < high; j++ {
		if list[j] <= pivot {
			list[i], list[j] = list[j], list[i]
			i ++
		}
	}

	list[i], list[high] = list[high], list[i]
	//--------------]]

	LomutoSort(list, low, i-1)
	LomutoSort(list, i+1, high)
}

/*
	霍尔快排：
	递归模板同老毛桃快排，返回低值区的锚点下标。不同的是霍尔快排以最左为基点比较元素，
	并从两端同时进行比较双向排序, 将双方都需调整位置的两个元素相互交换直到双方相交，
	然后将最左元素与右端当前锚点交换，最终实现基准点归位。

	评价——

	关键点——

 */
func HoareSort(list []int, low int, high int) {
	if low >= high {return}

	//[[----递归模板区
	
	pivot := list[low]
	right := high
	left := low
	for {
		for list[right] >= pivot && right > low{
			right --
		}
		for list[left] <= pivot && left < high{
			left ++
		}
		
		
		if left < right {
			list[left], list[right] = list[right], list[left]
		} else {
			break
		}
	}
	list[low], list[right] = list[right], list[low]
	//--------------]]
	

	HoareSort(list, low, right)

	HoareSort(list, right+1, high)
}