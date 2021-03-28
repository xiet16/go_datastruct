package ArraySorts

import (
	"fmt"
)

func CountSort(arr []int) []int {
	length:=len(arr)
	if length<=1 {
		return arr
	}

	maxValue := GetMaxValue(arr)
	countArr:=make([]int,maxValue+1)
	res:=make([]int,length)

	for i:=0;i<length;i++ {
		countArr[arr[i]]++  //统计次数
	}
	fmt.Println("统计次数:",countArr)

	for i:=1;i<=maxValue;i++ {
		countArr[i]+=countArr[i-1] //叠加
	}
	fmt.Println("次数叠加:",countArr )

    //展开数据
	for _,v:=range arr{
		res[countArr[v]-1] =v
		countArr[v]--
	}

	return res
}