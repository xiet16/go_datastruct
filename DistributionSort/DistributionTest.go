package DistributionSort

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func DistributionTest()  {
	thradMerge()
}

func syncMapTest(){
	smap = SyncMap{make(map[string]string),new(sync.RWMutex)}
	done =make(chan bool,2)

	go Write()
	go Write2()

	for {
		Read()
		if len(done) ==2{
			fmt.Println("done",smap.mymap)
			break
		}else {
			time.Sleep(1*time.Second)
		}
	}
}

func goroutineShellSort()  {
	arr:=make([]int,0)
	for i:=0;i<10000;i++ {
		arr=append(arr,i)
	}

	arr=GoroutineShellSort(arr)

	fmt.Println(arr)
}

func crateMergeFile()  {
	var filename string ="E:\\Go\\go_data_struct\\Files\\data.in"
	var count = 100000
	file,err:=os.Create(filename)
	if err!=nil {
		panic(err)
	}
	defer file.Close()

	piple:=RandomSource(count) //创建随机数组
	writer:=bufio.NewWriter(file) // 写入
	WriterLink(writer,piple)
	writer.Flush()  //刷新

	file,err=os.Open(filename)
	if err!=nil {
		panic(err)
	}

	defer file.Close()

    readPiple :=ReaderSource(bufio.NewReader(file),-1)
    counter := 0
	for v:=range readPiple {
		fmt.Println(v)
		counter++
		if counter >1000{
			break
		}
	}
}

func muiltiMerge()  {
	piple := Merge(InMemorySort(ArraySoure(4,5,1,2,8,3,9)), 
		           InMemorySort(ArraySoure(6,2,4,5,7,8,10,35)))

	for v:=range piple {
		fmt.Println(v)
	}
}

func thradMerge()  {
	p:=createPipleline("E:\\Go\\go_data_struct\\Files\\data.in",800000,4)
	writeToFile(p,"E:\\Go\\go_data_struct\\Files\\data.out")
	showFile("E:\\Go\\go_data_struct\\Files\\data.out")

	time.Sleep(20 *time.Second)
}