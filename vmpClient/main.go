
package main

import(

// "github.com/gin-gonic/gin"
  "net"
  "fmt"
  "bytes"
  "encoding/binary"
)

type proto_head struct{
	text_len int32
	bin_len int32
}

func main() {
	conn, err := net.Dial("tcp", "192.168.4.205:9030")
	if err!=nil{
		fmt.Println(err)
	}
	
	head:=&proto_head{}

	str := `{ "Message" : "LoginRequest", "Params" : { "Account" : "account", "Password" : "md5password" } }`
	head.text_len = int32(len(str))
	head.bin_len = 0
	

	databufio := new(bytes.Buffer)
	err = binary.Write(databufio,binary.LittleEndian,head)
	fmt.Println(err == nil)
	
	var baffer0 bytes.Buffer

	baffer:=databufio.Bytes()
	baffer1:=[]byte(str)
	baffer0.Write(baffer)
	baffer0.Write(baffer1)	
	fmt.Println(baffer0.Bytes())
	go  conn.Write(baffer0.Bytes())

	var ch = make(chan int)
	<-ch
}