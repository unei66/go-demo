package httprpc
import (
	"testing"
	"httprpc/server"
	"time"
	"httprpc/client"
)

func Test_httprpc(t *testing.T){
	go server.Serv()
	time.Sleep(3*time.Second)

	go client.Client()

	time.Sleep(5*time.Second)
}