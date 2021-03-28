package ArraySorts

import "strings"

func SelectSortMax(arr []int) int {
	length := len(arr)
	if length<=1 {
		return arr[0]
	}

	max :=0
	for i:=0;i<length;i++ {
		if arr[i]>max {
			max =arr[i]
		}
	}

	return max
}

//原理，每一次循环，都能取出最小值或者最大值，并将他们放到前面的位置
func SelectSort() []int {
	arr:=[]int{6,2,8,7,1,0,5}
	length := len(arr)
	if length<=1 {
		return arr
	}

	for i:=0;i<length-1;i++ {
		for j:=i+1;j<length;j++ {
		if arr[i]<arr[j] {
			arr[i],arr[j]=arr[j],arr[i]
		}
	}
}

	return arr
}

func StringSelectSort(arr []string) []string {
	length :=len(arr)
	if length<=1 {
		return arr
	}
	for i:=0;i<length-1;i++ {
		for j:=0;j<length;j++ {
			if strings.Compare(arr[i],arr[j]) >0 {
				arr[i],arr[j] =arr[j],arr[i]
			}
		}
	}

	return arr
}