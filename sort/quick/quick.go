/*
When given an array, `quicksort()` splits it up into three parts based on a "pivot" variable. Here, the pivot is taken to be the element in the middle of the array (later on you'll see other ways to choose the pivot).

All the elements less than the pivot go into a new array called `less`. All the elements equal to the pivot go into the `equal` array. And you guessed it, all elements greater than the pivot go into the third array, `greater`. This is why the generic type `T` must be `Comparable`, so we can compare the elements with `<`, `==`, and `>`.

Once we have these three arrays, `quicksort()` recursively sorts the `less` array and the `right` array, then glues those sorted subarrays back together with the `equal` array to get the final result. 
*/
package quick

// 简单快排
func Sort(list []int) []int{

	if len(list) <= 1 {
		return list
	}


	pivot := list[len(list)/2]
	//基点分割函数
	partitionFunc :=  func (arr []int, pivot int) ([]int, []int, []int){
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

	copy(list, Sort(less))
	copy(list[len(less):], equal)
	copy(list[(len(less)+len(equal)):], Sort(greater))

	return list
}

/*
  Lomuto's partitioning algorithm.

  This is conceptually simpler than Hoare's original scheme but less efficient.

  The return value is the index of the pivot element in the new array. The left
  partition is [low...p-1]; the right partition is [p+1...high], where p is the
  return value.

  The left partition includes all values smaller than or equal to the pivot, so
  if the pivot value occurs more than once, its duplicates will be found in the
  left partition.
*/
func LomutoSort(list []int, low int, high int){
	if low >= high {return}
	// We always use the highest item as the pivot.
	pivot := list[high]

	// This loop partitions the array into four (possibly empty) regions:
	//   [low  ...      i] contains all values <= pivot,
	//   [i+1  ...    j-1] contains all values > pivot,
	//   [j    ... high-1] are values we haven't looked at yet,
	//   [high           ] is the pivot value.
	i := low
	for j := low+1; j < high; j++ {
		if list[j] <= pivot {
			list[i], list[j] = list[j], list[i]
			i ++ 
		}
	}

	list[i], list[high] = list[high], list[i]

	LomutoSort(list, low, i-1)
	LomutoSort(list, i+1, high)
}