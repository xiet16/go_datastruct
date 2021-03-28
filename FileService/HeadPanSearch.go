package FileService

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//硬盘搜索
func HeadPanSerch()  {
	fi,err := os.Open("E:\\Go\\go_data_struct\\Files\\csdn_password.txt")
	if err!=nil {
		return
	}

	rd:=bufio.NewReader(fi)

	t1:=time.Now()
	for {
		line,_,err:=rd.ReadLine()
		if err == io.EOF{
			break
		}
		linecont:=string(line)
		if strings.Contains(linecont,"xiet")  {
			fmt.Println(linecont)
		}
	}
	interval := time.Since(t1)
	fmt.Println("headpan used time:",interval)
}