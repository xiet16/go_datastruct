package ArraySorts

import "fmt"

/*
插入排序法
{6,2,8,7,1,0,5}
原理：
第一步：取出要对比的数据
第二步：找到要插入的位置
第三步：插入位置往前或后移
第四步：插入第一步取出的数据

区别选择排序法
选择排序法是分别循环找第一个位置，第二个位置，第三个位置，第n个位置的值
插入排序是给值找位置
*/

func InsertSort(arr []int) []int {
	length :=len(arr)
	if length <=1{
		return arr
	}

	for i:=1;i<length;i++ {
       j:=i-1
       backup:=arr[i]  //取出要找位置的值
		for j>=0 &&arr[j]<backup {
			arr[j+1]=arr[j]  //往后移动
			j--
		}
		arr[j+1]=backup //插入
		fmt.Println(arr)
	}

	return arr
}