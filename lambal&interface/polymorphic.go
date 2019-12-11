package main

import "fmt"

type animal interface {
	run()
	age() int
}

type cat struct{
}

func (_ *cat)run(){
	fmt.Println("cat run!")
}

func (_ *cat)age() (x1 int){
	x1 = 10
	return 
}

type dog struct{

}

func (_ *dog) run() {
	fmt.Println("dog run")
}
func(_ *dog)age() (int){
	return 20
}

func factory(name string) animal{
	switch name{
	case "dog":
			var m *dog = new(dog)
			return m
	case "cat":
			//return &cat{};
			return new(cat)
		default:
			panic("no animal!");
	}
}

func main(){
	a:=factory("cat")
	a.run()
	fmt.Print(a.age(),"\n")
}

