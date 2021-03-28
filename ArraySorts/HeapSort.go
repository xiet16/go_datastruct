package ArraySorts

import "fmt"

func HeapSort(arr []int) []int  {
	length:=len(arr)
	for i:=0;i<length;i++ {
		lastLength:=length-i
		HeapSortMax(arr,lastLength)
		if i<length {
			arr[0],arr[lastLength-1]=arr[lastLength-1],arr[0]
		}

	}
	fmt.Println(arr)
	return arr
}

//堆排序获取最大值
func HeapSortMax(arr []int,length int) []int {
	if length <=1{
      return arr
	}

	deap :=length/2+1  //计算深度
	for i:=deap-1;i>=0;i-- {
		maxValue := i
		leftChildren:=2*i +1  //左孩子节点
		rightChildren :=2*i+2  //右孩子节点
		//防止数组溢出
		if leftChildren <= length-1 && arr[leftChildren]<arr[maxValue]{
			//arr[leftChildren],arr[maxValue]=arr[maxValue],arr[leftChildren]
			//这样写也行
			maxValue = leftChildren
		}

		if rightChildren <=length-1 && arr[rightChildren]<arr[maxValue] {
			//arr[rightChildren],arr[maxValue]=arr[maxValue],arr[rightChildren]
			maxValue =rightChildren //
		}

		if maxValue!=i {
			arr[maxValue],arr[i]=arr[i],arr[maxValue]
		}
		fmt.Println(arr)
	}
	return arr
}