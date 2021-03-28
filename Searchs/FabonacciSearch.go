package Searchs

import "fmt"

/*
斐波那契查找:
二分查找和中值查找的扩展，因为它的比例逐渐接近于黄金分割点，也叫黄金分割点查找
*/

//构建斐波那契数组
func makeFabArray(arr []int) []int {
	length:=len(arr)
	max:= arr[length-1]
	fabArr:=make([]int,3,3)
	fabArr[0]=1
	fabArr[1]=2
	fabArr[2]=3
	for i:=2;i <length-1 && fabArr[i]< max;i++ {
	   	fabArr=append(fabArr,fabArr[i]+fabArr[i-1])
	}
	return fabArr
}

func FabonacciSearch(arr []int,value int) int {
	length:=len(arr)
	//获取斐波那契数组
	fabArr:=makeFabArray(arr)
	fmt.Println(fabArr)

	//创建填充的数组
	fillArrLength:=fabArr[len(fabArr)-1]
	fillArr :=make([]int,fillArrLength)
	for i:=0;i<length;i++ {
		fillArr[i]=arr[i]
	}

	for i:=length;i<fillArrLength;i++ {
		fillArr[i]=arr[length-1]
	}
	fmt.Println(fillArr)

	left,mid,right:=0,0,length-1
	index := len(fabArr)-1
	for left<right {
		mid=left+fabArr[index-1]-1  //斐波那契切割
		if value < fillArr[mid]{
			right=mid-1
			index--
		}else if value>fillArr[mid] {
            left=mid+1
            index -=2  //为什么减2，减1 肯定是不行，但为什么是2
		}else {
			if mid>right {
				return right  //越界
			}else {
				return mid
			}
		}
	}
    return -1
}