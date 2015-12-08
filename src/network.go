package main
import (
	"net"
	"fmt"
	"os"
	"bytes"
	"io"
)

func networkTest(){
	icmp()
}

func icmp(){
	target:="www.baidu.com"

	conn,err:=net.Dial("ip4:icmp",target)
	checkErr(err)

	var msg [512]byte
	msg[0]=8
	msg[1]=0
	msg[2]=0
	msg[3]=0
	msg[4]=0
	msg[5]=13
	msg[7]=37
	len:=8
	check:=checkSum(msg[0:len])
	msg[2]=byte(check>>8)
	msg[3]=byte(check&255)

	_,err=conn.Write(msg[0:len])
	checkErr(err)

	_,err=conn.Read(msg[0:])
	checkErr(err)

	fmt.Println("Got response")

	if msg[5]==13 {
		fmt.Println("identifider matches")
	}

	if msg[7]==37 {
		fmt.Println("sequence matches")
	}

	os.Exit(0)
}

func checkErr(err error)  {
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
		os.Exit(1)
	}
}

func checkSum(msg[]byte)uint16{
	sum:=0

	for n:=1;n<len(msg)-1;n+=2{
		sum+=int(msg[n])*256+int(msg[n+1])
	}

	sum=(sum>>16)+(sum&0xffff)
	sum+=(sum>>16)
	var answer uint16=uint16(^sum)
	return answer
}

func readFully(conn net.Conn)([]byte,error){
	defer  conn.Close()

	result:=bytes.NewBuffer(nil)

	var buf [512]byte
	for{
		n,err:=conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err!=nil{
			if err==io.EOF{
				break
			}
			return nil,err
		}
	}

	return result.Bytes(),nil
}
