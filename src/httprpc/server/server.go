package server
import (
	"net/rpc"
	"net"
	"log"
	"net/http"
)

type Echo int

func (t *Echo) Hi (args string,reply *string)error{
	*reply="server echo:"+args
	return nil;
}

func Serv(){
	rpc.Register(new(Echo))
	rpc.HandleHTTP()
	l,e:=net.Listen("tcp",":1234")
	if e!=nil{
		log.Fatal("listen error:",e)
	}

	http.Serve(l,nil)
}