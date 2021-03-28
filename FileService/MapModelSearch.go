package FileService

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//模型排序
func MapModelSearch()  {
	const  N=6428633 //内存
	mapArr :=make(map[string]string,N)

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
	 	mapArr[strs[0]] =strs[1]
		i++
	}

	for {
		fmt.Println("请输入要查询的用户名")
		var name string
		fmt.Scanf("%s",&name)

		t1:=time.Now()
		v,ok:=mapArr[name]
		if ok{
			fmt.Println(name,"---",v)
		}
		fmt.Println("serchname ",name,"used ",time.Since(t1))
	}
}