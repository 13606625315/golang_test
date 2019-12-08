package main

import "fmt"

type chen interface {
	test(int, int) int
}

type xiexie struct {
}

func (_ xiexie) test(a int, b int) int {
	println("33333")
	return a + b
}

func test_1() chen { //接口作为返回值
	var a xiexie
	return a	
}


func main() {
	app := func(a, b int) func(int, int) int {
		fmt.Println("11111")
		return func(a, b int) int {
			fmt.Println("222222")
			return a + b
		}

	}
	fmt.Println("test1")	
	var bbb xiexie
	var aaa chen
	aaa = bbb
	aaa.test(1, 2)
	fmt.Println("test2")	
	bpp := app(1, 2)
	bpp(1, 2)

	fmt.Println("test3")	
	aaa = test_1()
	aaa.test(3,2)
}
