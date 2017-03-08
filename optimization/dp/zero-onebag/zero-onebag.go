/*
zero-one bag 0-1背包问题

状态转移方程为d(i, j)=max{ d(i-1, j), d(i-1, j-V[i-1]) + W[i-1] }。
其中状态d(i-1, j)表示前i个宝石装到体积为j的背包里能达到的最大价值，V[i]和w[i]分别表示第i个宝石对应的体积和的价值。

迭代实现过程中的状态转移方程是运算过程的抽象，是种数学表示，故实现过程中的逻辑有时会不太显然。
但只要数学过程是正确的，则可保证运算过程的正确性
 */
package zeroonebag

func max(first , second int) int{
	if first > second {return first}

	return second
}
func run(stones []map[string]int, capacity int) int{
	
	return recurse(stones, capacity)
	// return iterate(stones, capacity)
}
/*
迭代法
 */
func iterate(stones []map[string]int, capacity int)  int{
	bagCValues := []int{}
	//初始化状态,注意足够容量！——i <= capacity
	for i:=0;i<=capacity;i++{
		bagCValues = append(bagCValues, 0)
	}

	for _,mapitem := range stones{//遍历宝石
		/*
		 * 遍历容量空间，要注意的是capacityLast初始化值时capacity而非capacity-1
		 * capacityLast语义是剩余容量，剩余容量可能为零也可能为最大容量！
		 */
	    for capacityLast := capacity; capacityLast >= mapitem["weight"]; capacityLast--{
	    	
    		CurrentValue := bagCValues[capacityLast - mapitem["weight"]] + mapitem["value"]
    		
			bagCValues[capacityLast] = max(bagCValues[capacityLast], CurrentValue)
	    }
	}
	return bagCValues[len(bagCValues)-1]
}
/*
递归法
 */
func recurse(stones []map[string]int, capacity int)  int{
	length := len(stones)
	if length==0 || capacity==0 {
        return 0 
	}else{  
	    for stoneIndex := length-1;stoneIndex>=0;stoneIndex-- {  
	        if stones[stoneIndex]["weight"] > capacity{  
	            return recurse(stones[:length-1],capacity) 
	        }else{  
	            temp1:=recurse(stones[:length-1],capacity) 
	            temp2:=recurse(stones[:length-1],capacity-stones[stoneIndex]["weight"])+stones[stoneIndex]["value"] 
	            return max(temp2, temp1)
	        }  
	    }  
    }
    return 0
}