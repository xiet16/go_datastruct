package ArraySorts

import (
	"fmt"
)

/*
基数排序:

*/

func RadixSort(arr []int) []int {
   maxValue:=GetMaxValue(arr)

	for bit:=1; maxValue/bit>0 ;bit*=10{
		arr=BitShort(arr,bit)
	}
	return arr
}

//寻找最大值

func GetMaxValue(arr []int) int {
	length :=len(arr)
	if length<=1 {
		return arr[0]
	}

	maxValue :=0
	for i:=0;i<length;i++ {
		if arr[i] > maxValue{
			maxValue=arr[i]
		}
	}
	return maxValue
}

func BitShort(arr []int,bit int) []int {
   length:=len(arr)
   bitCounts :=make([]int,10)
	for i:=0;i<length;i++ {
		num :=(arr[i]/bit)%10 //分层 10,100,1000,1w
		bitCounts[num]++  //统计余数相等的个数
	}
	fmt.Println(bitCounts)

	for i:=1;i<10;i++ {
		bitCounts[i]+=bitCounts[i-1]  //叠加计算位置
	}

	tmp:=make([]int,10)
	for i:=length-1;i>=0;i-- {
		num:=(arr[i]/bit)%10
		tmp[bitCounts[num]-1]=arr[i]//计算排序的位置
		bitCounts[num]--
	}

	for i:=0;i<length;i++{
		arr[i]=tmp[i]
	}
	return arr
}