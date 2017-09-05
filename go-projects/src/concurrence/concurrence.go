package concurrence

import(
	"sync"
	"errors"
	"time"
)

//尽量不要使用这种共享数据，因为需要加锁费时，要用channel代替
var data chan Info = make(chan Info, 8*10000);
var queueLock sync.Mutex;
var num int = 0;

type Info struct{
	Len  int
	Head string
	Body string
}

type myQueue interface{
	Put() error
	Get(info Info) error
	GetNum() int;
	CloseChan()bool;
}


func (info Info) Put() (err error){
	if info.Len == 0 {
		return errors.New("put info len is null");
	}
	data<-info;
	queueLock.Lock();
	num++;
	queueLock.Unlock();
	time.Sleep(1000);
	return nil;
}


func (info Info) Get() (err error){
	info = <-data;
	queueLock.Lock();
	num--;
	queueLock.Unlock();
	if info.Len < 0 {
		return errors.New("get info len < 0");
	}

	return nil;
}

func (info Info) GetNum() int{
	return num;
}

//channel closing principe
func (info Info)CloseChan() bool{
	close(data);
	_, ok := <-data;
	return ok;
}