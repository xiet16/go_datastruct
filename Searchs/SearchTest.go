package Searchs

import "fmt"

func SearchTest()  {
	arr:=[]int{}
	for i:=0;i<10;i++ {
		arr = append(arr,i)
	}

   index:= FabonacciSearch(arr,9)
	if index >0{
		fmt.Println(index," value: ", arr[index])
	}else {
		fmt.Println("未找到数据")
	}

}

func middleTest(arr []int)  {
	indes:= MiddleSearch(arr,89)
	for _,v:= range indes {
		fmt.Println("find arr value:",arr[v])
	}
}