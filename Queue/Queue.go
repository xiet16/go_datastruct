package Queue

/*
非线程安全队列
*/

type IQueue interface {
	Size() int //大小
	Front() interface{}
	End() interface{}
	IsEmpty() bool
	EnQueue(data interface{}) //入队
	DeQueue()interface{}
	Clear()//清空
}

type Queue struct {
	dataStroe []interface{} //
	theSize int
}

func NewQueue()*Queue  {
   queue :=new(Queue)
   queue.dataStroe =make([]interface{},0)
   queue.theSize=0
   return queue
}

func (q *Queue)Size() int {
   return q.theSize
}
func (q *Queue)Front() interface{}{
	if q.theSize ==0 {
		return nil
	}
	return q.dataStroe[0]
}
func (q *Queue)End() interface{}{
	if q.theSize==0 {
		return nil
	}
	return q.dataStroe[q.theSize-1]
}
func (q *Queue)IsEmpty() bool{
  return q.theSize ==0
}
func (q *Queue)EnQueue(data interface{}) {
  q.dataStroe =append(q.dataStroe,data)
  q.theSize ++
}
func (q *Queue)DeQueue()interface{}{
	if q.theSize==0 {
		return nil
	}
	data:=q.dataStroe[0]
	if q.theSize>1 {
		q.dataStroe = q.dataStroe[1:q.theSize] //这里很浪费资源
	}else {
		q.dataStroe =make([]interface{},0)
	}
	q.theSize--
	return data
}

func (q *Queue)Clear() {
  q.dataStroe = make([]interface{},0)
  q.theSize =0
}
