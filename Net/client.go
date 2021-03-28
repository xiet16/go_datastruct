package Net

import (
	"fmt"
	"net"
)

/*
等待服务器发送数据
客户端接收
内存排序
返回数据给服务器
*/

func ClientStart()  {
	go NetListenInit("127.0.0.1:7001")
	go NetListenInit("127.0.0.1:7002")
	//ClientSend("",[]int{})
}

func ClientSend(addr string,arr []int) {
	serverAddr,err :=net.ResolveTCPAddr("tcp4",addr)
	if err!=nil {
		panic(err)
	}

	conn,err:=net.DialTCP("tcp",nil,serverAddr) //连接服务器
	if err!=nil {
		panic(err)
	}

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

func ClientSendTest()  {
	serverAddr,err :=net.ResolveTCPAddr("tcp4","127.0.0.1:8848")
	if err!=nil {
		panic(err)
	}

	conn,err:=net.DialTCP("tcp",nil,serverAddr) //连接服务器
	if err!=nil {
		panic(err)
	}

	go func() {
		for {
			var inputstr string
			fmt.Scanln(&inputstr)
			_,err := conn.Write([]byte(inputstr))
			if err!=nil {
				fmt.Println( err.Error())
			}
		}
	}()


	go func() {
		for {
			buf:=make([]byte,1024)
			n,_:=conn.Read(buf)
			if n>0 {
				fmt.Println("receive server respons: ",string(buf[:n]))
			}
		}
	}()
}