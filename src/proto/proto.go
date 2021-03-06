package proto
import (
	"log"
	"github.com/golang/protobuf/proto"
)

func Parse(){
	test:=&Test{
		Label:proto.String("hello"),
		Type:proto.Int32(17),
		Optionalgroup:&Test_OptionalGroup{
			RequiredField:proto.String("good bye")
		},
	}

	data,err:=proto.Marshal(test)
	if err!=nil{
		log.Fatal("unmarshalling error:",err)
	}

	newTest:=&Test{}

	err=proto.Unmarshal(data,newTest)
	if err!=nil{
		log.Fatal("unmarshalling error:",err)
	}

	if(test.GetLabel()!=newTest.GetLabel()){
		log.Fatalf("data mismatch %q!=%q",test.GetLabel(),newTest.GetLabel())
	}
}