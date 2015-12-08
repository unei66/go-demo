package etcd
import (
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"time"
	"log"
)

func Conn(){
	cfg:=client.Config{
		Endpoints:[]string{"http://127.0.0.1:2379/","http://127.0.0.1:4001"},
		Transport:client.DefaultTransport,
		HeaderTimeoutPerRequest:time.Second,
	}

	c,err:=client.New(cfg)

	if err!=nil{
		log.Fatal(err)
	}

	kapi:=client.NewKeysAPI(c)

	log.Print("Setting '/foo' key with 'bar' value")
	resp,err:=kapi.Set(context.Background(),"/foo","bar",nil)
	if err!=nil{
		log.Fatal(err)
	}else{
		log.Printf("Set is done.Metadata is %q\n",resp)
	}

	log.Print("Getting '/foo' key value")
	resp,err=kapi.Get(context.Background(),"/foo",nil)
	if err!=nil{
		log.Fatal(err)
	}else{
		log.Printf("Get is done,Metadata is %q\n",resp)
		log.Printf("%q key has %q value\n",resp.Node.Key,resp.Node.Value)
	}

}