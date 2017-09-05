package main

import(
	//"basic"
	"concurrence"
	"fmt"
	"runtime"
	"time"
) 


func testChannel(){

	startTime := time.Now();
	
	numCpu := runtime.NumCPU();
	runtime.GOMAXPROCS(numCpu);
	fmt.Println("CPU num:", numCpu);
	
	var info concurrence.Info = concurrence.Info{Len:8, Head:"TCP", Body:"im ok"};
	for i:=0; i<1000; i++ {
		go info.Put();
	}

	//其实这个时间这样统计也是错误的	
	elapsed := time.Since(startTime);	
	fmt.Println("put data to channel waste time:", elapsed);
	/*
	num := info.GetNum();
	fmt.Println("info in Queue num is:", num);
	info.Get();
	fmt.Println("info->len:", info.Len,
				"info->head:", info.Head,
				"info->body", info.Body);
	*/
	//保证上面的协程都处理完了，go没有wait这种等待线程处理完的函数，需要自己管理
	time.Sleep(1000*1000*1000);

	num := info.GetNum();
	fmt.Println("info in Queue num is:", num);
	
}


func testSyntax(){
	// basic.TestVar();

	// basic.TestString();

	// basic.TestSclice();

	// basic.TestMap();

	//basic.TestInterface();

	//testChannel();
}

func main(){




}

