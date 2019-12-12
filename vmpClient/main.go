
package main

import(

// "github.com/gin-gonic/gin"
  "net"
  "fmt"
  "bytes"
  "encoding/binary"
  "log"
  "net/http"
  "context"
  "io"
  "time"
)

type proto_head struct{
	text_len int32
	bin_len int32
}

var conn net.Conn

var ch chan string

func send2Cloud(w http.ResponseWriter,p []byte){

	head:=&proto_head{}
	//str := `{ "Message" : "LoginRequest", "Params" : { "Account" : "XxWTug0UtcayU8sUF4YVnQ==", "Password" : "lingsuo" } }`
	str:= string(p)
	head.text_len = int32(len(str))
	head.bin_len = 0
	
	databufio := new(bytes.Buffer)
	err:= binary.Write(databufio,binary.BigEndian,head)
	fmt.Println(err == nil)
	
	var baffer0 bytes.Buffer
	baffer:=databufio.Bytes()
	baffer1:=[]byte(str)
	baffer0.Write(baffer)
	baffer0.Write(baffer1)	
	conn.Write(baffer0.Bytes())

	select {
    case res := <-ch:
		io.WriteString(w, res)
    case <-time.After(time.Second * 5):
		io.WriteString(w, "timeout")
    }

}

func handler(w http.ResponseWriter, r *http.Request) {
    //获取内容的长度
    length := r.ContentLength
    //创建一个字节切片
    body := make([]byte, length)
    //读取请求体
    r.Body.Read(body)
	log.Println( "请求体中的内容是：", string(body))
	send2Cloud(w,body)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	ch = make(chan string)
	var err error 
	conn, err = net.Dial("tcp", "192.168.4.205:9030")
	if err!=nil{
		fmt.Println(err)
	}	

	go func() {
		ctx := context.Background()
		for {
			txt, bin, err := readMessage(conn)
			if err != nil {
				log.Println("conn closed!", err)
				break
			}
			err = OnMessage(ctx, txt, bin)
			if err != nil {
				log.Println("on message error", err)
			}
		}
	}()

	http.HandleFunc("/getBody", handler)	
	if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}


}

func OnMessage(ctx context.Context, txt []byte, bin []byte) error {
	log.Println(string(txt))
	ch <- string(txt)
	return nil;
}