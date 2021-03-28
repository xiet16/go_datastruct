package Link

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

/*
将文件读入链表
*/

func FileStoreByLink(){
	fi,err:= os.Open("E:\\Go\\go_data_struct\\Files\\csdn_net.sql")
	if err!=nil {
		fmt.Println("打开文件错误:" +err.Error())
		return
	}

	defer fi.Close()
	rd:=bufio.NewReader(fi)
	linkList := NewSingleLinkList()
	for {
		line,_,err:=rd.ReadLine()
		if err==io.EOF {
			break
		}
		linestr :=string(line)
		//strs:=strings.Split(linestr," # ")
		node:=&SingleLinkNode{value: linestr}
		linkList.InsertNodeFront(node)
	}

	for {
		fmt.Println("请输入要查询的用户名")
		var name string
		fmt.Scanf("%s",&name)

		t1:=time.Now()
        fmt.Println(linkList.FindString(name))
		fmt.Println("serchname ",name,"used ",time.Since(t1))
	}
}