package bubblesort

import (
	"testing"
	"fmt"
)

func TestBubbleSort1(t *testing.T){
	values:=[]int{5,4,3,2,1}

	BubbleSort(values)

	for _,v:=range values{
		fmt.Println(v)
	}
}