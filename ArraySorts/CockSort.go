package ArraySorts

import (
	"fmt"
)

/*
鸡尾酒排序:
本质就是循环双边冒泡，每次冒泡都会有一个值沉底或值飘起
*/

func CockSort(arr []int)[]int  {
	length:=len(arr)
	if length<=1 {
		return arr
	}

	for i:=0;i<length;i++ {
		left :=0;
		right :=length-1
		for left<=right {
			if arr[left]<arr[left+1] {
				arr[left],arr[left+1]=arr[left+1],arr[left]
			}
			left++

			if arr[right]>arr[right-1] {
				arr[right],arr[right-1]=arr[right-1],arr[right]
			}
			right--
		}
		fmt.Println(i,arr)
	}
	return arr
}