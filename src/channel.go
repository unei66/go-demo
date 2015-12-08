package main
import (
	"fmt"
	"time"
	"sync"
)

func Count(ch chan int){
	ch<- 1
	fmt.Println("Counting")
}

func channelTest(){
	twoprint()
}

func simple(){
	chs:=make([]chan int,10)

	for i:=0;i<10;i++{
		chs[i]=make(chan int)
		go Count(chs[i])
	}

	for _,ch:=range(chs){
		a:=<-ch
		fmt.Println(a)
	}
}

func deadFor(){
	ch :=make(chan int,1)
	for{
		select{
		case ch<-0:
		case ch<-1:
		}

		i:=<-ch
		fmt.Println("value:",i)
	}
}

func timeout() {

	ch:=make(chan int,1)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(10*time.Second)
		timeout<-true
	}()

	select{
	case <-ch:
	case <-timeout:
		fmt.Println("timeout")
	}
}

var a string
var once sync.Once

func setup(){
	a="hello world"
}

func doprint(){
	once.Do(setup)
	fmt.Println(a)
}

func twoprint(){
	go doprint()
	go doprint()
}

