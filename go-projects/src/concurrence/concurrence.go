package concurrence

import(
	"fmt"
)

type myQueue interface{
	put() error
	get() error
}

type Info struct{
	len  int
	head string
	body string
}


var Data chan Info = make(chan Info, 1024);

func (info Info) put() error{
	Data<-info;
	defer func() {
		if err := recover();err != nil{
			fmt.Println(err);
		}
	}()
}


func (info Info) get() error{
	
}
