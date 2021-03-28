package FileService

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func PasswordSort()  {
	const  N=6428633 //内存
	strArr :=make([]string,N,N)

	fi,err := os.Open("E:\\Go\\go_data_struct\\Files\\csdn_password.txt")
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
		strArr[i]=string(line)
		i++
	}
    fmt.Println("sort start ")
	t1:=time.Now()
	strArr = QuickSortString(strArr)
	intervalTime := time.Since(t1)
	fmt.Println("sort end ",intervalTime)

	savePath:="E:\\Go\\go_data_struct\\Files\\password_sort.txt"
	saveFile,_:=os.Create(savePath)
	defer saveFile.Close()
	wr:=bufio.NewWriter(saveFile)
	for i:=0;i<len(strArr);i++ {
		fmt.Fprintln(wr,strArr[i])
	}

	fmt.Println(intervalTime)
}

func QuickSortString(arr []string)[]string{
	length:= len(arr)
	if length<=1 {
      return arr
	}

	n:=length/2
	base:=arr[n]

	low:=[]string{}
	high:=[]string{}
	mid :=[]string{}

	for i:=0;i<length;i++ {
		if base>arr[i] {
			low=append(low,arr[i])
		}else if base <arr[i] {
			high =append(high,arr[i])
		}else{
			mid =append(mid,arr[i])
		}
	}

	low,high =QuickSortString(low),QuickSortString(high)

	res:=append(append(low,mid...),high...)
    return res
}