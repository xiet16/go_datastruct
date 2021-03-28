package ArraySorts

/*
奇偶排序：
以奇数和偶数为基准，对整个数组进行排序，排序的元素只有两个元素（基数和右侧相邻元素）
如何证明奇偶排序的可行性？后续证明，这个博主有写(https://blog.csdn.net/lemon_tree12138/article/details/50605563)
*/

func OddEvenSort(arr []int) []int {
    length :=len(arr)
	if length <=1{
		return arr
	}
	n:=1
	isexchage := false
	for i:=n%2;i<length;i+=2 {
		if i+1<length && arr[i]<arr[i+1] {
			arr[i],arr[i+1]=arr[i+1],arr[i]
			isexchage =true
		}
	}
	n++
	for i:=n%2;i<length;i+=2 {
		if i+1<length && arr[i]<arr[i+1] {
			arr[i],arr[i+1]=arr[i+1],arr[i]
			isexchage =true
		}
	}

	if isexchage {
		arr = OddEvenSort(arr)
	}

	return arr
}