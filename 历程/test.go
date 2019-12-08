package main
import (
    "sync"
	"time"
)

var wg sync.WaitGroup

func say1(ch chan int) {
	for index := 0; index < 5; index++ {
		println("Hello %v",index)				
		ch<-index
	
	}
}

func say2(ch chan int) {
	for index := 0; index < 5; index++ {
	
		println("world %v",index)	
		<-ch
	
	}
}


func main() {

	var ch = make(chan int)
    go say1(ch)
    go say2(ch)
	
	time.Sleep(1e9)
}

