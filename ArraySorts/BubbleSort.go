package ArraySorts

import "fmt"

/*
冒泡排序：
循环两两交换，每一次内循环，都有会产生一个最大或者最小值，
最大最小值会沉底，其他值往上


6,2,8,7,1,0,5

第一步：第一次内循环
6和2对比，6,2,8,7,1,0,5
2和8对比，6,8,2,7,1,0,5
2和7对比，6,8,7,2,1,0,5
2和1对比，6,8,7,2,1,0,5
1和0对比，6,8,7,2,1,0,5
0和5对比，6,8,7,2,1,5,0
第二步：第二次内循环
6和8对比，8,6,7,2,1,5,0
6和7对比，8,7,6,2,1,5,0
6和2对比，8,7,6,2,1,5,0
2和1对比，8,7,6,2,1,5,0
1和5对比，8,7,6,2,5,1,0
依次类推
*/

func BubbleSortTest(arr []int) []int {
	length:=len(arr)
	if length<=1 {
       return arr
	}

	for i:=0;i<length-1;i++{
		isexchange:=false
		for j:=0;j<length-1-i;j++ {
			if arr[j]<arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
				isexchange =true
			}
		}
		if !isexchange {
			break
		}
		fmt.Println(arr)
	}
	return arr
}