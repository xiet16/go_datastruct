package DistributionSort

import (
	"fmt"
	"github.com/xiet16/go_data_struct/ArraySorts"
	"runtime"
	"sync"
)

/*
希尔排序是插入排序的一种，是插入排序的优化
希尔排序时间复杂度是 O(n^(1.3-2))，
空间复杂度为常数阶 O(1)。
希尔排序没有时间复杂度为 O(n(logn)) 的快速排序算法快 ，
因此对中等大小规模表现良好，
但对规模非常大的数据排序不是最优选择，总之比一般 O(n^2 ) 复杂度的算法快得多。
这是插入排序:
[6,2,8,7,1,0,5] 原数组
[6 2 8 7 1 0 5] 第一次循环从下标a1->a0 (前面的数据是循环后的结果)
[8 6 2 7 1 0 5] 第三次循环从下标a2-a0
[8 7 6 2 1 0 5]
[8 7 6 2 1 0 5]
[8 7 6 2 1 0 5]
[8 7 6 5 2 1 0]

这是插入排序
[7 2 8 6 1 0 5]
[7 2 8 6 1 0 5]
3  步长->start: 0  插入排序后: [7 2 8 6 1 0 5]  备注: 前面两个是子插入过程，这一行是结果，实际排序 a[0],a[3],a[6] 3 个值，需要排2轮
[7 2 8 6 1 0 5]
3  步长->start: 1  插入排序后: [7 2 8 6 1 0 5]  备注: 前面1行是子插入过程，这一行是结果，实际排序 a[1],a[4]
[7 2 8 6 1 0 5]
3  步长->start: 2  插入排序后: [7 2 8 6 1 0 5]  备注: 前面1行是子插入过程，这一行是结果，实际排序 a[2],a[5]
[7 2 8 6 1 0 5]
[8 7 2 6 1 0 5]
[8 7 6 2 1 0 5]
[8 7 6 2 1 0 5]
[8 7 6 2 1 0 5]
[8 7 6 5 2 1 0]
1  步长->start: 0  插入排序后: [8 7 6 5 2 1 0]  备注: 因为此时补偿为1，所以这个是一轮全插入

同一步长，是可以异步的
*/


func GoroutineShellSort(arr []int) []int{
	length:=len(arr)
	if length <2 {
		return arr
	}
	wg:=sync.WaitGroup{} //等待线程返回
	goroutinenum:=runtime.NumCPU() //系统cpu 倍数

	//压缩步长
	for gap:=length/2;gap>0;gap/=2 {
		wg.Add(goroutinenum) //需要等待的协程数
		ch := make(chan int,10000)  //通道

		//往通道传递参数
		go func() {
			for k:=0;k<gap;k++ {
				ch <- k
			}

			close(ch) //关闭管道
		}()

		//开启协程，不断的接收通道信息
		//原理就是，用一个协程给channel 传递参数， 用n个协程去执行任务
		for k:=0;k<goroutinenum;k++ {
			go func() {
				for v:=range ch{
					ArraySorts.ShellStep(arr,v,gap)
				}

				wg.Done() //完成
			}()
		}
		wg.Wait()

		fmt.Println("current gap:",gap)
	}

    return arr
}