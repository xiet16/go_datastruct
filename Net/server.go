package Net

import (
	"net"
)

/*
服务器得到数据
切割数组
发送给客户端
等待客户端结果
合并排序
*/



func ServerStart() {
	//开启服务器
    go ServerController()
}



//发送原始数据
func ServerSend()  {
	serverAddr,err :=net.ResolveTCPAddr("tcp4","127.0.0.1:8848")
	if err!=nil {
		panic(err)
	}

	conn,err:=net.DialTCP("tcp",nil,serverAddr) //连接服务器
	if err!=nil {
		panic(err)
	}

	arr:=[]int{1,5,6,2,8,7,1,2,5,9,7}
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
		data :=IntToByte(-1)
		data = append(data,IntToByte(arr[i])...)
		conn.Write(data)
	}
	
	end:=IntToByte(-1)
	end = append(end,IntToByte(1)...)

}

