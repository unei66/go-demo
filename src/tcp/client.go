package tcp
import (
	"net"
	"fmt"
//	"os"
	"time"
//	"bufio"
)

type Clienter struct {
	client net.Conn
	isAlive bool
	SendStr chan string
	RecvStr chan string
}

func (c *Clienter)Connect()bool{
	if c.isAlive{
		return true
	}

	var err error
	c.client,err=net.Dial("tcp","localhost:8088")
	if err!=nil{
		fmt.Printf("Failure to connect:%s\n",err.Error())
		return false
	}

	c.isAlive=true

	return true
}

func (c *Clienter) Echo(){
	line:=<-c.SendStr
	c.client.Write([]byte(line))
	buf:=make([]byte,1024)
	n,err:=c.client.Read(buf)
	if err!=nil{
		c.RecvStr<- string("Server close....")
		c.client.Close()
		c.isAlive=false
		return
	}

	time.Sleep(1*time.Second)
	c.RecvStr<-string(buf[0:n])
}

func Work(tc *Clienter){
	if !tc.isAlive{
		if tc.Connect(){
			tc.Echo()
		}else{
			<-tc.SendStr
			tc.RecvStr<-string("server close...")
		}

	}else{
		tc.Echo()
	}
}

func Client(){
	var tc Clienter
	tc.SendStr=make(chan string)
	tc.RecvStr=make(chan string)
	if !tc.Connect(){
		return
	}

//	r:=bufio.NewReader(os.Stdin)

	for i:=0;i<100;i++{
		line:="heha"
		go Work(&tc)
		tc.SendStr<-line
		s:=<-tc.RecvStr
		fmt.Printf("back:%s\n",s)
	}
}