package tcp
import (
	"net"
	"fmt"
	"os"
	"time"
)

const(
MAX_CONN_NUM=5
)

func EchoFunc(conn net.Conn){
	defer conn.Close()

	buf:=make([]byte,1024)
	for{
		n,err:=conn.Read(buf)
		if err!=nil{
			return
		}

		_,err=conn.Write(buf[0:n])
		if err!=nil{
			return
		}
	}
}

func Serve(){
	listener,err:=net.Listen("tcp","127.0.0.1:8088")
	if err!=nil{
		fmt.Println("error listening",err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("running.....")

	var curConnNum int=0	//当前链接总数
	connChan:=make(chan net.Conn)
	chConnChange:=make(chan int)	//统计当前连接数channel

	go func(){
		for connChange:=range chConnChange{
			curConnNum+=connChange
		}
	}()

	go func(){
		for _=range time.Tick(1*time.Second){
			fmt.Printf("cur conn num:%d\n",curConnNum)
		}
	}()

	for i:=0;i<MAX_CONN_NUM;i++{
		go func(){
			for conn:=range connChan{
				chConnChange<- 1
				EchoFunc(conn)
				chConnChange<- -1
			}
		}()
	}

	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("Error Accept:",err.Error())
			return
		}

		connChan<- conn

	}
}