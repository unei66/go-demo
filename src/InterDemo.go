package main
import "fmt"

type Person struct{
	name string
	age int
}

func (p Person)printMsg(){
	fmt.Printf("I am %s,and my age is %d.\n",p.name,p.age)
}

func (p Person) eat(s string){
	fmt.Printf("%s is eating %s...\n",p.name,s)
}

func (p Person) drink(s string){
	fmt.Printf("%s is drinking %s...\n",p.name,s)
}

type PeopleDrink interface{
	drink(s string)
}

type PeopleEat interface {
	eat(s string)
}

type PeopleEatDrink interface{
	eat(s string)
	drink(s string)
}

type People interface{
	printMsg()
	PeopleDrink
	PeopleEat
}

type Foodie struct{
	name string
}

func (f Foodie)eat(s string){
	fmt.Printf("I am foodie,%s,My favorite food is %s\n",f.name,s)
}

func test(){
	var p1 People
	p1=Person{"Rain",23}
	p1.printMsg()
	p1.drink("orange juice")

	var p2 PeopleEat
	p2=Person{"Sun",24}
	p2.eat("chaffy dish")

	var p3 PeopleEat
	p3=Foodie{"James"}
	p3.eat("noodle")

	p3=p1
	p3.eat("noodle")

	if p4,ok:=p2.(People);ok{
		p4.drink("water")
		fmt.Println(p4)
	}
}
