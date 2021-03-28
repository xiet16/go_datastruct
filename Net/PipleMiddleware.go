package Net

import (
	"bufio"
	"github.com/xiet16/go_data_struct/DistributionSort"
	"net"
	"os"
	"strconv"
)

func creatNetworkPiple(filename string,filesize int,chunkCount int) <-chan int {
	//var filename string ="E:\\Go\\go_data_struct\\Files\\data.in"
	var count = 100000
	file,err:=os.Create(filename)
	if err!=nil {
		panic(err)
	}
	defer file.Close()

	piple:=DistributionSort.RandomSource(count) //创建随机数组
	writer:=bufio.NewWriter(file) // 写入
	DistributionSort.WriterLink(writer,piple)
	writer.Flush()  //刷新

    chunksize := filesize/chunkCount //大小
    sortArr := []string{} // 分布式的地址集合

    DistributionSort.Init()  //初始化
    file,err=os.Open(filename)
	if err!=nil {
		panic(err)
	}


	for i:=0;i<chunkCount;i++ {
		file.Seek(int64(i*chunksize),0) //移动文件指针位置
		source := DistributionSort.ReaderSource(bufio.NewReader(file),chunksize) //读取
		addr:=":"+strconv.Itoa(7000+i) //开辟地址
		NetWrite(addr,DistributionSort.InMemorySort(source)) //写入到分布式主机
		sortArr =append(sortArr,addr) //地址复制
	}

    sortResults :=[]<-chan int{}
    for _,addr:=range sortArr{
    	sortResults = append(sortResults,NetRead(addr))
	}

    return DistributionSort.MergeN(sortResults...)
}

//往tcp 服务发送数据数据
func NetWrite(addr string,in <-chan int) {
	listen,err:=net.Listen("tcp",addr)
	if err!=nil {
		panic(err)
	}

	go func() {
		defer listen.Close() //关闭网络

	  	conn,err := listen.Accept() //接收信息
		if err!=nil {
			panic(err)
		}
		defer conn.Close()

	  	writer := bufio.NewWriter(conn)
	  	defer writer.Flush()

        DistributionSort.WriterLink(writer,in)
	}()
}

func NetRead(addr string) <-chan int{
	out:=make(chan int,1024)

    go func() {
		conn,err := net.Dial("tcp",addr)
		if err!=nil {
			panic(err)
		}

		r:=DistributionSort.ReaderSource(bufio.NewReader(conn),-1)

		for v:=range r {
			out <- v
		}

		close(out)
	}()
	return out
}