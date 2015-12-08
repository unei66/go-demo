package client
import (
	"net/rpc"
	"log"
	"fmt"
)

func Client(){
	client,err:=rpc.DialHTTP("tcp","127.0.0.1:1234")
	if err!=nil{
		log.Fatal("dialing:",err)
	}

	var args="hello pc"
	var reply string
	err=client.Call("Echo.Hi",args,&reply)
	if err!=nil{
		log.Fatal("arith err:",err)
	}

	fmt.Println("client receive:",reply)
}