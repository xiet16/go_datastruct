package StackArray

import (
	"fmt"
	"io/ioutil"
)

//递归主要是循环调用子类，还有存储
func GetAllFile(path string,files []string) ([]string,error) {
   rd,err:=ioutil.ReadDir(path)
	if err !=nil{
		return files,err
	}
	for _,fi:=range rd {
		if fi.IsDir() {
           dir:=path+"\\" + fi.Name()
           files = append(files,dir) //
           files ,_ =GetAllFile(dir,files) //文件夹递归处理
		} else {
			files = append(files,path+"\\"+fi.Name())
		}
	}
	return files,nil
}

func FileFuncRecTest()  {
	files,err:=GetAllFile("E:\\分布式学习",make([]string,0))
	if err !=nil{
		fmt.Println(err)
	}

	for _,f :=range files{
		fmt.Println(f)
	}
}

//使用栈
func FileStackRecTest(){
	files,err:=GetAllFileByStack("E:\\分布式学习",make([]string,0))
	if err !=nil{
		fmt.Println(err)
	}

	for _,f :=range files{
		fmt.Println(f)
	}
}

func GetAllFileByStack(path string,files []string)([]string,error){
	stack:=NewStack()
	stack.Push(path)
	if !stack.IsEmpty() {
	 	dir:= stack.Pop().(string)
		rd,err:=ioutil.ReadDir(dir)
		if err !=nil{
			return files,err
		}
		for _,fi:=range rd {
			if fi.IsDir() {
				dir:=path+"\\" + fi.Name()
				files = append(files,dir) //
				files ,_ =GetAllFileByStack(dir,files) //文件夹递归处理
			} else {
				files = append(files,path+"\\"+fi.Name())
			}
		}
	}
	return files,nil
}