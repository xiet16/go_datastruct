package ArraySorts

import "fmt"

/*
希尔算法:
步长如何定义:
步长如何调整，以及如何停止
*/

func ShellSort(arr []int)[]int{
   length:=len(arr)
	if length<=1 {
		return arr
	}
	gap :=length/2
	for gap>0 {
		for i:=0;i<gap;i++ {
			ShellStep(arr,i,gap)
			fmt.Println(gap," 步长->start:",i," 插入排序后:",arr)
		}
		gap/=2
	}
	return arr
}

func ShellStep(arr []int,start int,gap int) []int {
	length:=len(arr)
	for i:=start+gap;i<length;i+=gap {
		//间隔为gap 的插入排序
		backup :=arr[i]
		j:=i-gap
		//数据循环往后移动
		for j>=0 && backup>arr[j]{
            arr[j+gap] = arr[j]
            j-=gap
		}
		arr[j+gap]=backup //插入
		//fmt.Println(arr)
	}
	return arr
}