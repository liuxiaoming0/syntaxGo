package basic

import "fmt"

func MyPrint(a int, b int) (int){
	d := a + b;
	return d;
}

func TestVar(){
	
	var a ,b = 10, 20;
	c := MyPrint(a, b);
	fmt.Println(a,"+", b,"=",c)

	const d string = "go";
	const
	(e = iota
	 f
	 g)

	fmt.Println(d, e, f, g);

	var h int = 5;
	// s := 'a';
	var s byte = 'a';
	s = (byte)(h);
	fmt.Println(h, s);
}

func TestString(){

	var a string = "go编程";
	l := len(a);
	fmt.Println(len(a));
	for i := 0; i < l; i++ {
		fmt.Println(a[i]);
	} 

	for i, ch := range a{
		fmt.Println(i, ch);
	}

}

func TestSclice()(){

	var a [5]int = [5]int{1, 2, 3, 4, 5};
	var b []int = a[:2];
	for i, v := range b{
		fmt.Println(i, v);
	}

	var c[]int = []int{10, 20, 30};
	c = append(c, 40, 50, 60);
	c = append(c, b...);
	for i, v := range c{
		fmt.Println(i, v);
	}
}

func TestMap(){
	var a map[int]string = map[int]string{1: "hello", 2: "world"};
	a[3] = "go";
	a[4] = "gogo";
	for i, v := range a{
		fmt.Println(i, v);
	}
}

/*******************interface*******************/
type animal interface{
	print() int
	feature() string
}

type People struct{
	name 	string
	age  	int
	f	string
}

type Pig struct{
	weight 	int
	female 	bool
	where   string
	f	string
}

func (p *People) print()int{
	fmt.Println(p.name, p.age, p.f);
	return 0;
}

func (p People) feature()string{
	return p.f;
}


func (p Pig) print()int{
	fmt.Println(p.weight, p.female, p.f);
	return 0;
}

func (p Pig) feature()string{
	return p.f;
}

func TestInterface(){

	var p People = People{name:"xiaohong", age:22, f:"is a people"};
	var p1 Pig = Pig{weight:100, female:true, where:"in pigpen ", f:"is a pig"};
	p.print(); // p is a interface
	p1.print();
	
}