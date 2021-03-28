package FileService

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const N = 8000000

func MemorySearch(){
	fi,err := os.Open("E:\\Go\\go_data_struct\\Files\\csdn_password.txt")
	if err!=nil {
		return
	}

	rd:=bufio.NewReader(fi)
    strArr :=make([]string,N,N)

	for {
		line,_,err:=rd.ReadLine()
		if err == io.EOF{
			break
		}
		strArr = append(strArr,string(line))
	}

	t1:=time.Now()
	for i:=0;i<len(strArr)-1;i++ {
		if strings.Contains(strArr[i],"xiet")  {
			fmt.Println(strArr[i])
		}
	}
	interval := time.Since(t1)
	fmt.Println("used time:",interval)
}