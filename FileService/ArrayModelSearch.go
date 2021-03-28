package FileService

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type CsdnModel struct {
	userName string
	userPwd string
	userEmail string
}

//模型排序
func ArrayModelSearch()  {
	const  N=6428633 //内存
	modelArr :=make([]CsdnModel,N,N)

	fi,err := os.Open("E:\\Go\\go_data_struct\\Files\\csdn_net.sql")
	if err!=nil {
		fmt.Println("open file error :",err.Error())
		return
	}

	defer fi.Close()

	rd:=bufio.NewReader(fi)
	i :=0
	for  {
		line,_,err:=rd.ReadLine()
		if err == io.EOF {
			break
		}
		linestr :=string(line)
		strs:=strings.Split(linestr," # ")
		modelArr[i] =CsdnModel{
			userName:strs[0],
			userPwd: strs[1],
			userEmail: strs[2],
		}
		i++
	}
    fmt.Println("file count:",i)
	t1:=time.Now()
	modelArr = QuickSortStruct(modelArr)
	fmt.Println("排序所需时间: ",time.Since(t1))

	for  {
		fmt.Println("请输入要查询的用户名")
		var name string
		fmt.Scanf("%s",&name)

		t1=time.Now()
		for j:=0;j<N;j++ {
			if modelArr[j].userName == name {
				fmt.Println(modelArr[j].userName,"---",modelArr[j].userEmail)
			}
		}
		fmt.Println("scan searchname ",name,"used ",time.Since(t1))

		t1 =time.Now()
		indexArr:= BinarySearchStruct(modelArr,CsdnModel{userName: name})
		for _,v:=  range indexArr {
			fmt.Println("binary serchname ",modelArr[v],"used ",time.Since(t1))
		}
	}
}


func BinarySearchStruct(arr[] CsdnModel ,data CsdnModel) []int {
	left:=0
	right:=len(arr)-1
	indes:=[]int{}
	for left <right{
		mid := (left+right)/2
		if arr[mid].userName>data.userName {
			left =mid+1
		}else if arr[mid].userName<data.userName{
			right=mid-1
		}else {
			indes=append(indes,mid)
			for i:=mid+1;i<len(arr);i++ {
				if arr[i].userName == data.userName {
					indes=append(indes,i)
				}else {
					break
				}
			}

			for i:=mid-1;i>0;i-- {
				if arr[i].userName == data.userName {
					indes=append(indes,i)
				}else {
					break
				}
			}
			return indes
		}
	}
	return indes
}


func QuickSortStruct(arr []CsdnModel) []CsdnModel {
	length:=len(arr)
	if length <=1 {
		return arr
	}

	base:=arr[0]
	low:=[]CsdnModel{}
	high:=[]CsdnModel{}
	mid:=[]CsdnModel{}

	mid=append(mid,base)
	for i:=1;i<length;i++ {
		if base.userName>arr[i].userName {
			low=append(low,arr[i])
		}else if base.userName <arr[i].userName {
			high=append(high,arr[i])
		}else {
			mid=append(mid,arr[i])
		}
	}
	low,high = QuickSortStruct(low),QuickSortStruct(high)
	res:=append(append(high,mid...),low...)
	return res
}