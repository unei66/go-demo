package jsonrpc
import (
	"testing"
	"time"
)


func Test_jsonrpc(t *testing.T){
	go Serv()
	time.Sleep(2*time.Second)

	go Client()
	time.Sleep(2*time.Second)
}
