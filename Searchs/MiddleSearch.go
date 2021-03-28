package Searchs

/*
中值搜索:
根据值比例求中值
1,2,3,4,5,6,7,8,9,10 共10位数，取 8 在哪个位置
left:0
right:9
data/就得出值
(arr[right]-arr[left])/(right-left) 没份的间隔是多少
求data 大概在第几份 (data-arr[left])*(right-left)/(arr[right]-arr[left])  = n
中间值下标就是 left+ n
*/

func MiddleSearch(arr []int,data int) []int {
	length:=len(arr)
	if length <1 {
		return nil
	}

	left:=0
	right:=length-1
    res:=[]int{}
	if left<right {
		//mid:=(right+left)/2    //优化欠
		allv := float64(arr[right]-arr[left])
		leftv := float64(data-arr[left])
		diff := float64(right-left)
		mid:= int(float64(left)+ leftv/allv*diff)

		if mid<0 || mid >length {
			return nil
		}
		if arr[mid] >data {
			right =mid-1
		}else if arr[mid]<data {
			left = mid+1
		}else {
          res =append(res,mid)
			for i:=mid+1;i<=right;i++ {
				if arr[i]==data {
					res=append(res,i)
				}else {
					break
				}
			}
			for i:=mid-1;i>=left;i-- {
				if arr[i]==data {
					res=append(res,i)
				}else {
					break
				}
			}
		}
	}

	return res
}