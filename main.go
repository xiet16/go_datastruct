package main

import (
   "flag"
   "github.com/xiet16/go_data_struct/DistributionSort"
   "github.com/xiet16/go_data_struct/Net"
   "log"
   "os"
   "runtime"
   "runtime/pprof"
)

var (
   cpuprofile = flag.String("cpuprofile","","write cpu profile to file")
   memprofile = flag.String("memprofile","","write mem profile to file ")
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

   flag.Parse()
   if *cpuprofile !="" {
      f,err := os.Create(*cpuprofile)
      if err!=nil {
         log.Fatal(err)
      }else {

      }

      pprof.StartCPUProfile(f)
      defer  pprof.StopCPUProfile()
   }

   go func() {
      Net.ServerStart()
      Net.ClientStart()
   }()

   DistributionSort.DistributionTest()
   if *memprofile !="" {
      f,err:=os.Create(*memprofile)
      if err!=nil {
         log.Fatal("could not create memory profile: ", err)
      }
      runtime.GC()

      if err:=pprof.WriteHeapProfile(f);err!=nil {
         log.Fatal("could not write memory profile: ", err)
      }
      f.Close()
   }
}