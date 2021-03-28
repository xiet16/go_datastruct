package StackArray

import (
	"fmt"
	"io/ioutil"
)

func GetLevelFile(path string,files []string,level int) ([]string,error) {
	levstr:=""
    for i:=1;i<level;i++{
		levstr +="|--"
	}
	rd,err:=ioutil.ReadDir(path)
	if err !=nil{
		return files,err
	}
	for _,fi:=range rd {
		if fi.IsDir() {
			dir:=path+"\\" + fi.Name()
			files = append(files,dir) //
			files ,_ =GetLevelFile(dir,files,level+1) //文件夹递归处理
		} else {
			files = append(files,path+"\\"+fi.Name())
		}
	}
	return files,nil
}

//使用栈
func LevelFileStackRecTest(){
	files,err:=GetLevelFile("E:\\分布式学习",make([]string,0),1)
	if err !=nil{
		fmt.Println(err)
	}

	for _,f :=range files{
		fmt.Println(f)
	}
}

func GeLevelFileByStack(path string,files []string)([]string,error){
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