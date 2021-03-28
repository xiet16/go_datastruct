package Searchs

//二分查找法
func MyBinarySearch(arr[] int ,data int) []int {
	 length:=len(arr)
	if length ==0 {
		return nil
	}
	 left:=0
	 right :=length -1
	 mid := length/2

	 res :=[]int{} //结果下标集合
	if arr[mid]<data{
		res = MyBinarySearch(arr[0:mid],data)
	}else if arr[mid]>data{
		res = MyBinarySearch(arr[mid:],data)
	}else{
		 res =append(res,mid)
		for i:=mid+1;i<=right;i++ {
			if arr[i]==data{
				res=append(res)
			}
		}

		for j:=mid-1;j>=left;j-- {
			if arr[j]==data{
				res=append(res)
			}
		}
	}
	return res
}

func BinarySearch(arr[] int ,data int) int {
     left:=0
     right:=len(arr)-1
	for left <right{
		mid := (left+right)/2
		if arr[mid]>data {
			left =mid+1
		}else if arr[mid]<data{
			right=mid-1
		}else {
		    return mid
		}
	}
	return -1
}