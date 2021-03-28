package DistributionSort

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var starttime time.Time

func Init()  {
	starttime =time.Now()
}

func UseTime()  {
	fmt.Println(time.Since(starttime)) //统计小号时间
}

//内存排序
func InMemorySort(in <-chan  int)<-chan int {
	out := make(chan int,1024) //创建新管道

	//构建线程
	go func() {
		data :=[]int{}  //创建数组，存储数据并排序
		for v:=range in{
			data =append(data,v)
		}
		fmt.Println("数据读取完成 ", time.Since(starttime))

		//排序
        sort.Ints(data)
		for _,v:=range data{
			out<-v  //压入数据
		}
		close(out) //关闭管道
	}()

	return out
}

//合并
func Merge(in1,in2 <-chan int) <-chan int {
  out :=make(chan int,1024)
  
  go func() {
  	v1,ok1 := <-in1
  	v2,ok2 := <-in2

	  for ok1||ok2 {
		  if !ok2 || (ok1 && v1<=v2) {
			  out <-v1  //取出v1，压入，再次读取
			  v1,ok1 =<-in1
		  }else{
		  	out <-v2
		  	v2,ok2 =<-in2
		  }
	  }
	  close(out)
  }()
  return out
}


//读取

func ReaderSource(reader io.Reader,size int) <-chan int{
	out:=make(chan int,1024)
	go func() {
		buf :=make([]byte,8)
		readsize :=0
		for {
			n,err:=reader.Read(buf)
		    readsize +=n
			if n>0 {
				out <- int(binary.BigEndian.Uint64(buf)) //数据压入
			}

			if err!=nil ||(size!=-1&&readsize>=size){
				break  //跳出循环
			}
		}
		close(out)
	}()

	return out
}

//写入
func WriterLink(writer io.Writer,in <-chan int)  {
	for v:=range in {
		buf := make([]byte,8)
		binary.BigEndian.PutUint64(buf,uint64(v)) //字节转换
		writer.Write(buf)
	}
}


//随机数数组

func RandomSource(count int) <-chan int {
	out :=make(chan int,1024)

	go func() {
		for i:=0;i<count;i++ {
			out <-rand.Int()
		}
		close(out)
	}()

	return out
}

//多路合并

func MergeN(inputs ... <-chan int) <-chan int {
	if len(inputs) ==1 {
		return inputs[0]
	}else {
		m:=len(inputs)/2
		return Merge(MergeN(inputs[:m]...),MergeN(inputs[m:]...)) //递归
	}
}

func ArraySoure(nums... int)<-chan int{
	out :=make(chan int,1024)
	go func() {
		for i:=0;i<len(nums); i++{
			out <- nums[i]
		}
		close(out)
	}()
	return out
}
