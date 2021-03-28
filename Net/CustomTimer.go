package Net

import (
	"log"
	"net"
	"time"
)

/*
自定义定时器
*/


//超过30 秒 退出
func HeartBeat(conn net.Conn,heartChan chan byte,timeout int)  {
	select {
	case hc:= <-heartChan:
		//中断超时时间，计数器归零
		log.Println("heartbeat receive: ",string(hc))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout)*time.Second))
	case <-time.After(time.Second*30):
		log.Println("time out ") //客户端超时
		conn.Close()
	}
}

func HearChanHadler(n []byte,beatch chan byte)  {
	for _,v :=range n{
		beatch <- v
	}
	close(beatch) //关闭管道
}