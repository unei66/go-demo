
package main

import (
	"fmt"
	"math"
	"strings"
	"os"
	"time"
	"io"
	"strconv"
)
var c,java,python bool

type Stringer interface{
	String() string
}

func add(x int,y int)int{
	return x+y;
}

func swap(x,y string)(string,string){
	return y,x;
}

func pow(x,n,lim float64)float64{
	if v:=math.Pow(x,n);v<lim{
		return v
	}else{

	}

	return lim
}

func main(){
//	fmt.Println(add(12,13))
//	a:="haha";
//	b:="gogo";
//	a,b=swap(a,b)
//	fmt.Println(a,b)
//	fmt.Println(c)
//
//	sum:=0
//	for i:=0;i<10;i++{
//		sum+=i
//	}
//
//	fmt.Println(sum)
//
//	fmt.Println(pow(2,2,10))
//
//	myDefer()
//
//	myDeferStatck()
//
//	point()
//
//	structTest()
//
//	array()
//
//	slice()
//
//	makeSlice()
//
//	rangePrint()
//
//	mapPrint()
//
//	wordCount()
//
//	adderTest()
//
//	MethodReceiver()
//
//	TestInterface()
//
//	testError()
//
//	stringReader()
//
//	goTest()

//	sumTest()

//	testFib()

//	testClosure()111

//	aa :=CC{121111}
//	printInterface(aa)
//
//	var bb Stringer=new(CC)
//	printInterface(bb.String())
//	test()
	networkTest()
}

type CC struct{
	a int
}

func (cc *CC) String() string{
	return strconv.Itoa(cc.a)
}

func printInterface(args ... interface{}){

	for _,arg:=range args{

		switch arg.(type){
		default:
			if v,ok:=arg.(Stringer);ok{
				val:=v.String();
				fmt.Println(val)
			}
		}
	}
}

func myDefer(){
	defer fmt.Println("world")
	fmt.Println("hello")
	fmt.Println("  ")
}

func myDeferStatck(){
	fmt.Println("counting")
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}

func point(){
	fmt.Println("point...................")
	i,j:=42,2701
	p:=&i
	fmt.Println(*p)
	*p=21
	fmt.Println(i)

	p=&j
	*p=*p/37
	fmt.Println(j)
}

func structTest(){
	fmt.Println("structTest........................")
	type V struct{
		x int
		y int
	}

	fmt.Println(V{1,2})
}

func array(){
	fmt.Println("array...........................")
	var a [2]string
	a[0]="hello"
	fmt.Println(a)
}

func slice(){
	fmt.Println("slice.............................")
	s:=[]int{2,3,4,5,6,7,8,9,10}
	fmt.Println("s==",s)

	for i:=0;i<len(s);i++{
		fmt.Printf("s[%d]=%d\n",i,s[i])
	}

	fmt.Println("s[1:4]==",s[1:4])

	fmt.Println("s[:3]==",s[:3])

	fmt.Println("s[4:]==",s[4:])
}

func makeSlice(){
	fmt.Println("makeSlice.............................")
	a:=make([]int,5)
	printSlice("a",a)
	b:=make([]int,0,5)
	printSlice("b",b)
	c:=b[:2]
	printSlice("c",c)
	d:=c[2:5]
	printSlice("d",d)
}

func printSlice(s string,x[]int){
	fmt.Printf("%s len=%d cap=%d %v \n",s,len(x),cap(x),x)
}

func rangePrint(){
	fmt.Println("rangePrint.................................")
	pow:=make([]int,10)
	for i:=range pow{
		pow[i]=1<<uint(i)
	}

	for _,value:=range pow{
		fmt.Printf("%d\n",value)
	}
}

func mapPrint(){
	fmt.Println("mapPrint......................................")
	type Ver struct{
		Lat,Long float64
	}

	var m map[string]Ver;
	m=make(map[string]Ver)
	m["Bell Labs"]=Ver{
		40.00,50.00,
	}

	fmt.Println(m["Bell Labs"])

	elm,ok:=m["Bell"];
	fmt.Println(ok,elm)

	elm,ok=m["Bell Labs"]
	fmt.Println(ok,elm)
}

func wordCount(){
	s:="hello world haha haha go go cc cc aa aa"
	str_arr:=strings.Fields(s)
	wordCountMap:=make(map[string]int)
	for _,value:=range str_arr{
		num,exist:=wordCountMap[value]
		if(exist){
			num=num+1
			wordCountMap[value]=num
		}else{
			wordCountMap[value]=1
		}
	}

	fmt.Println(wordCountMap)
}

func adder() func(int) int{
	sum:=0
	return func(x int)int{
		sum+=x
		return sum
	}
}

func adderTest(){
	fmt.Println("adderTest......................")
	pos,neg:=adder(),adder()

	for i:=0;i<10;i++{
		fmt.Println(pos(i),neg(-2*i))
	}
}

type Vertex struct{
	X,Y float64
}

func (v *Vertex) Scale(f float64){
	v.X=v.X*f
	v.Y=v.Y*f
}

func (v *Vertex)Abs() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func MethodReceiver(){
	fmt.Println("MethodReceiver...................")
	v:=&Vertex{3,4}

	fmt.Printf("Before scaling:%+v,Abs:%v\n",v,v.Abs())
	v.Scale(5)

	fmt.Println("After scaling:%+v,Abs:%v\n",v,v.Abs())
}

type Reader interface {
	Read(b []byte)(n int,err error)
}

type Writer interface {
	Write(b []byte)(n int,err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func TestInterface(){
	fmt.Println("TestInterface..................")
	var w Writer
	w=os.Stdout
	fmt.Fprintf(w,"hello world\n")
}

type MyError struct{
	When time.Time
	What string
}

func (e *MyError)Error()string{
	return fmt.Sprintf("at %v,%s",e.When,e.What)
}

func run() error{
	return &MyError{time.Now(),"it didn't work"}
}

func testError(){
	fmt.Println("testError.............................")
	if err:=run();err!=nil{
		fmt.Println(err)
	}
}

func stringReader(){
	fmt.Println("stringReader...........................")
	r:=strings.NewReader("hello,reader")

	b:=make([]byte,8)
	for{
		n,err:=r.Read(b)
		fmt.Printf("n=%v,err=%v,b=%v\n",n,err,b)
		fmt.Printf("b[:n]=%q\n",b[:n])
		if err==io.EOF{
			break
		}
	}
}

func say(s string){
	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func say2(s string){
	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func goTest(){
	fmt.Println("goTest..........................")
	go say("world")
	say2("hello")
}

func sum(a[] int,c chan int){
	sum:=0
	for _,v:=range a{
		sum+=v
	}

	c<-sum
}

func sumTest(){
	fmt.Println("sumTest....................")
	a:=[]int{7,2,8,-9,4,0}

	c:=make(chan int)
	go sum(a[:len(a)/2],c)

	go sum(a[len(a)/2:],c)

	x,y:=<-c,<-c

	fmt.Println(x,y,x+y)

}

func fibonacci(n int,c chan int){
	x,y:=0,1
	for i:=0;i<n;i++{
		c<-x
		x,y=y,x+y
	}

	close(c)
}

func testFib(){
	fmt.Println("testfib...................")
	c:=make(chan int,10)
	go fibonacci(20,c)
	for i:=range c{
		fmt.Println(i)
	}
}


