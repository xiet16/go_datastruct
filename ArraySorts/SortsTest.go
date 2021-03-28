package ArraySorts

import "fmt"

func SortsTest()  {
	//arr:=SelectSort()
    //arr:=StringSelectSort([]string{"aa","a1","d","b","i"})
	//arr:= InsertSort([]int{6,2,8,7,1,0,5})
	arr:=ShellSort([]int{6,2,8,7,1,0,5})
	//arr:=BubbleSortTest([]int{6,2,8,7,1,0,5})
	//arr:=HeapSort([]int{6,2,8,7,1,0,5})
    //arr:= QuickSortByRand([]int{6,2,8,7,16,10,34,78,90,8,9,18,16,25})
	//arr:=OddEvenSort([]int{6,2,8,7,16,10,34,78,90,8,9,18,16,25})
	//arr:=MyMergeSort([]int{6,2,8,7,16,10,34,78,90,8,9,18,16,25})
	//arr:=MergeSort([]int{6,2,8,7,16,10,34,78,90,8,9,18,16,25})
	//arr:=ShellSort([]int{6,2,8,7,16,10,34,78,90,8,9,18,16,25})
	//arr:=CountSort([]int {4,4,6,5,3,2,8,1,7,5,6,0,10})
	//arr:=RadixSort([]int {4,4,6,32,5,166,3,350,2,87,15,97,35,16,100,180})
	//arr:=RadixSort([]int {4,4,6,32,5,166,3,350,2,87})
	//arr:=CockSort([]int {4,4,6,32,5,166,3,350,2,87,15,97,35,16,100,180})
	//data :=[]int {4,4,6,32,5,166,3,350,2,87,15,97,35,16,100,180}
	//arr:= QuickInsertSort(data,0,len(data)-1)

	fmt.Println(arr)
	for _,v:=range arr{
		fmt.Println(v)
	}
}