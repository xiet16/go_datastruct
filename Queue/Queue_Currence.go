package Queue

import (
	"fmt"
	"io/ioutil"
)

/*
队列的特点:
先进先出
*/


//使用队列递归文件
func QueueFileCurrence(path string,files []string)([]string,error) {
   queue:=NewQueue()
	queue.EnQueue(path)
	for ;; {
		path:=queue.DeQueue()
		if path == nil {
			break
		}
		rd,err :=ioutil.ReadDir(path.(string))
		if err !=nil{
			return files,err
		}
		for _,fi:=range rd {
			if fi.IsDir() {
				fullpath:=path.(string)+"\\" + fi.Name()
				files = append(files,fullpath)
				fmt.Println("dir value: ",fullpath)
				queue.EnQueue(fullpath)
				//files ,_ =QueueFileCurrence(dir,files) //文件夹递归处理

			} else {
				fullpath:=path.(string)+"\\"+fi.Name()
				files = append(files,fullpath)
				fmt.Println("File:",fullpath)
			}
		}
	}

	return files,nil
}