package Queue

//自定义一个Int,用于测试
type MyInt int

func (x MyInt)Less(than Item) bool{
	return x <than.(MyInt)
}