package ArraySorts

/*
归并排序：
将数据切割成多端排序
*/

func MergeSort(arr []int) []int {
	length:=len(arr)
	if length<=1 {
		return arr
	}
	n:=length/2
	frontArr :=MergeSort(arr[:n])
	rearArr :=MergeSort(arr[n:])

	return merge(frontArr,rearArr)
}

func merge(leftArr,rightArr []int)[]int  {
	leftIndex:=0
	rightIndex :=0
	res:=[]int{}
	for leftIndex<len(leftArr)&&rightIndex<len(rightArr){
		if leftArr[leftIndex] >rightArr[rightIndex] {
			res=append(res,leftArr[leftIndex])
			leftIndex ++
		}else if leftArr[leftIndex] < rightArr[rightIndex] {
			res = append(res,rightArr[rightIndex])
			rightIndex++
		}else {
			res =append(res,leftArr[leftIndex],rightArr[rightIndex])
			leftIndex++
			rightIndex++
		}
	}

	for len(leftArr)>leftIndex {
		res =append(res,leftArr[leftIndex])
		leftIndex++
	}
	for len(rightArr)>rightIndex {
		res =append(res,rightArr[rightIndex])
		rightIndex++
	}
	return res
}

func MyMergeSort(arr []int) []int {
	length:=len(arr)
	if length<=1 {
      return arr
	}

	//n:=rand.Intn(length)
	n:=length/2
	frontArr :=arr[:n]
	rearArr :=arr[n:]
	res:=[]int{}
	//继续拆分排序
	frontArr = MyMergeSort(frontArr)
    rearArr = MyMergeSort(rearArr)

    //合并
	for i:=0;i<len(frontArr);i++ {
		data:=MyMergeData(frontArr[i],rearArr)
		if len(data)<1 {
			res=append(res,frontArr[i])
		}else {
			res=append(res,data...)
			res=append(res,frontArr[i])
			rearArr=rearArr[len(data):]
		}
	}
	if len(rearArr) >0{
		res =append(res,rearArr...)
	}
	return res
}

//
func MyMergeData(base int,arr []int) []int {
	res :=[]int{}
	for _,v:=range arr{
		if base<=v {
			res = append(res, v)
		}
	}
	return res
}

