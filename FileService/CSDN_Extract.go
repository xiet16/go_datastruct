package FileService

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

//提取csdn 文件数据
func CsdnExtract() error {
 fi,err:=os.Open("E:\\Go\\go_data_struct\\Files\\csdn_net.sql")
 if err!=nil{
	 return	errors.New("文件读取狮跑")
 }

 savePath :="E:\\Go\\go_data_struct\\Files\\csdn_password.txt"
 saveFile,err := os.Create(savePath)
 defer saveFile.Close()
 save := bufio.NewWriter(saveFile) //创建写入对象


 defer fi.Close()
	br:= bufio.NewReader(fi)
	lineCount :=0
	for {
		line,_,err:= br.ReadLine()
		if err==io.EOF {
			break
		}
		lineCount++
		linestr :=string(line)
		strs:=strings.Split(linestr," # ")
		fmt.Fprintln(save,strs[1]) //保存密码
	}
	fmt.Println(lineCount)
	save.Flush()
	return nil
}