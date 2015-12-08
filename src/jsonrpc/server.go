package jsonrpc
import (
	"net"
	"net/rpc"
	"log"
	"net/rpc/jsonrpc"
)

type Echo int

func (p Echo) Hi (args string,reply *string) error{
	*reply="jsonrpc echo:"+args
	return nil
}

func Serv() error{
	lis,err:=net.Listen("tcp",":1234")
	if err!=nil{
		return err
	}

	defer lis.Close()

	srv:=rpc.NewServer()
	if err:=srv.RegisterName("Echo",new(Echo));err!=nil{
		return err
	}

	for{
		conn,err:=lis.Accept()
		if err!=nil{
			log.Fatalf("lis.Accept:%v\n",err)
		}

		go srv.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}