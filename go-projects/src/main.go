package main

import(
	//"basic"
	"concurrence"
	"fmt"
	"runtime"
	"time"
) 


func testChannel(){

	var info concurrence.Info = concurrence.Info{Len:8, Head:"TCP", Body:"im ok"};
	for i:=0; i<1000000; i++ {
		go info.Put();
	}
	/*
	num := info.GetNum();
	fmt.Println("info in Queue num is:", num);
	info.Get();
	fmt.Println("info->len:", info.Len,
				"info->head:", info.Head,
				"info->body", info.Body);

	num = info.GetNum();
	fmt.Println("info in Queue num is:", num);
	*/
}

func main(){

	// basic.TestVar();

	// basic.TestString();

	// basic.TestSclice();

	// basic.TestMap();

	//basic.TestInterface();

	startTime := time.Now();

	numCpu := runtime.NumCPU();
	runtime.GOMAXPROCS(numCpu);
	fmt.Println("CPU num:", numCpu);


	testChannel();

	elapsed := time.Since(startTime);

	fmt.Println("put data to channel waste time:", elapsed);

}

