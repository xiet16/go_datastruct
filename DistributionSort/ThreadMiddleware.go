package DistributionSort

import (
   "bufio"
   "fmt"
   "os"
)

//多线程调用中间件
func createPipleline(filename string,filesize int,chunkCount int)<-chan int {
   chunkSize := filesize/chunkCount //数量
   sortResults := []<-chan int{} //爱旭结果，每个元素是一个掼蛋
   Init()
   for i:=0;i<chunkCount;i++ {
      file,err := os.Open(filename)
      if err!=nil {
         panic(err)
      }
      file.Seek(int64(1*chunkSize),0)  //跳转文件指针
      source :=ReaderSource(bufio.NewReader(file),chunkSize)
      sortResults =append(sortResults,InMemorySort(source)) //对结果排序
   }
   return MergeN(sortResults...)
}

//写入文件
func writeToFile(in <- chan int,filename string)  {
   file,err:=os.Create(filename)
   if err!=nil {
      fmt.Println(err.Error())
      panic(err)
   }
   defer file.Close()

   writer:=bufio.NewWriter(file) // 写入
   defer    writer.Flush()  //刷新
   WriterLink(writer,in)
}

//显示文件
func showFile(filename string)  {
   file,err:=os.Open(filename)
   if err!=nil {
      panic(err)
   }
   defer file.Close()

   p:=ReaderSource(bufio.NewReader(file),-1)
   counter := 0
   for v:=range p {
      fmt.Println(v)
      counter++
      if counter >1000{
         break
      }
   }
}