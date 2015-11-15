package main
import "fmt"


func testClosure(){
	fmt.Println("TestClosure...................")
	var j int=5

	a:=func() (func()) {
		var i int=10
		return func(){
			fmt.Printf("i,j:%d, %d\n",i,j)
		}
	}

	a()()

	j=10

	a()()
}