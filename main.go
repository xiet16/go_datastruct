package main

import (
   "github.com/xiet16/go_data_struct/Net"
   "os"
   "os/signal"
   "syscall"
)

func main()  {
   //StackArray.StackArrayTest()
   //ArrayList.IteratorTest()
   //ArrayList.StackArrayTest()
   //ArrayList.StackArrayIteratorTest()
   //StackArray.StackRecurrence()
   //StackArray.RecTest()
   //StackArray.FileStackRecTest()
   //Queue.QueueFileCurrence("E:\\分布式学习",make([]string,0))
   //Queue.CircleQueueTest()
   //Link.LinkStackTest()
   //ArraySorts.SortsTest()
   //FileService.HeadPanSerch()
   //FileService.MemorySearch()
   //FileService.SortAndSearch()
   //FileService.ArrayModelSearch()
   //Searchs.SearchTest()
   //Link.LinkStackTest()
   //Searchs.SearchTest()
   //Queue.PriorityQueueTest()
   //Hash.HashTableTest()
   //DistributionSort.DistributionTest()

   go func() {
      Net.ServerStart()
      Net.ClientStart()
   }()

   quit := make(chan os.Signal)
   signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
   <-quit
}