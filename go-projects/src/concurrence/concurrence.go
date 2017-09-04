package concurrence

import(
	"sync"
	"errors"
)

var data chan Info = make(chan Info, 1024);
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
	queueLock.Lock();
	data<-info;
	num++;
	queueLock.Unlock();
	return nil;
}


func (info Info) Get() (err error){
	queueLock.Lock();
	info = <-data;
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