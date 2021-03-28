package Net

import (
	"fmt"
	"github.com/xiet16/go_data_struct/DistributionSort"
	"net"
	"sort"
	"strconv"
)

//服务器监听
func ServerController() {
	//读取数据
    arrayList := [][]int{{1,12,3,5,7,9,3},{5,6,9,7,8,0,6}}
    outs:=make([]<-chan int,0)
    last :=make(<-chan int,1024)

    //调配任务，向客户端发送数据 并开启接收监听
	for i := 0; i < 2; i++ {
		clientAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:700"+strconv.Itoa(i+1))
		if err != nil {
			panic(err)
		}

		conn, err := net.DialTCP("tcp", nil, clientAddr)
		if err != nil {
			panic(err)
		}

		//发送原始数据
		NetSendData(arrayList[i],conn)

		//结果叠加
		//接收数据   全双工
		outs =append(outs,ServerMsgHandler(conn))
	}

	last = DistributionSort.MergeN(outs...)
	fmt.Println("服务端归并开始，长度为:",len(outs))
	for v:=range last{
		fmt.Printf("%d ", v)
	}
}

func NetSendData(arr []int,conn net.Conn) {

	length:=len(arr)
	//-1 -1 开始传输
	//1 1
	//1 5
	//1 6
	//... 每次发送一个数字
	//-1 1 结束

	start := IntToByte(-1)
	start = append(start,IntToByte(-1)...)
	conn.Write(start)

	for i:=0;i<length;i++ {
		data :=IntToByte(1)
		data = append(data,IntToByte(arr[i])...)
		conn.Write(data)
	}

	end:=IntToByte(-1)
	end = append(end,IntToByte(1)...)
	conn.Write(end)
}

//
func ServerMsgHandler(conn net.Conn) <-chan int{
	buf :=make([]byte,16)
	arr :=make([]int,0,0)
	out :=make(chan int ,1024)
	for {
		n,err:=conn.Read(buf)
		if err!=nil {
			fmt.Println("conn close")
			return nil
		}

		if n==16 {
			data1 := BytesToInt(buf[:len(buf)/2])
			data2 :=BytesToInt(buf[len(buf)/2:])

			if data1==-1 && data2 ==-1 {
				arr = make([]int,0,0) //开辟数组
			}else if data1==-1 && data2 ==1{
				for i:=0;i<len(arr);i++{
					out <- arr[i]
				}
				close(out)
				return out
			}else if data1 ==1{
				arr = append(arr, data2)
			}
		}
	}

	return nil
}


//带心跳的单节点测试
func MsgHandle(conn net.Conn)  {
	buf := make([]byte,16)
	defer conn.Close()
	arr := make([]int,0,0) //切片
	for {
		n,err:=conn.Read(buf)
		if err!=nil {
			fmt.Println("conn close")
			return
		}

		if n==16 {
			data1 := BytesToInt(buf[:len(buf)/2])
			data2 :=BytesToInt(buf[len(buf)/2:])

			if data1==-1 && data2 ==-1 {
				arr = make([]int,0,0) //开辟数组
			}else if data1==-1 && data2 ==1{
				fmt.Println("客户端接收到数据:",arr)     //传输完成
				sort.Ints(arr)
				fmt.Println("客户端排序后的数据:",arr)
				NetSendData(arr,conn)
			}else if data1 ==1{
				arr = append(arr, data2)
			}
		}

		//beatch:=make(chan byte)
		//go HeartBeat(conn,beatch,30)  //这里不用for 是为什么
		//go HearChanHadler(IntToByte(1),beatch)
	}
}


//初始化监听 单向通信测试
func NetListenInit(addr string)  {
	lister, err := net.Listen("tcp", addr)
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}

	defer lister.Close() //延迟关闭

	for {
		conn,err:=lister.Accept()
		if err!=nil {
			panic(err)
		}

		//开始接收
		go MsgHandle(conn)
	}
}