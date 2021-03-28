package Queue

import "errors"

const QueueSize = 100

type CircleQueue struct {
	data [QueueSize]interface{}
	front int //头部位置
	rear int //尾部位置
}

func InitQueue(q *CircleQueue)  {
	//头部和尾部重合为空
	q.front =0
	q.rear =0
}

func QueueLength(q *CircleQueue) int {
	return (q.rear-q.front+QueueSize)%QueueSize
}

func EnQueue(q *CircleQueue,data interface{}) error {
	if (q.rear+1)%QueueSize == q.front%QueueSize {
		return errors.New("队列已满")
	}

	q.data[q.rear]=data //入队
	q.rear = (q.rear+1)%QueueSize
	return nil
}

func DeQueue(q *CircleQueue)(data interface{},err error)  {
	if q.front == q.front {
		return nil, errors.New("队列为空")
	}

	res:=q.data[q.front] //取出数据
	q.data[q.front]=0 //清空数据
    q.front=(q.front+1)%QueueSize
	return res,nil
}