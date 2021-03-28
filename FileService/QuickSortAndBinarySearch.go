package FileService

import (
	"fmt"
	"github.com/xiet16/go_data_struct/Searchs"
	"github.com/xiet16/go_data_struct/ArraySorts"
)

func SortAndSearch()  {
	arr:= ArraySorts.QuickSort([]int {4,4,6,32,5,166,3,350,2,87,15,97,35,16,100,180})
	fmt.Println(arr)
	index := Searchs.BinarySearch(arr,15)
	if index >=0{
		fmt.Println(index," get data :", arr[index])
	}else {
		fmt.Println("未找到数据")
	}
}