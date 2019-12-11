package main
 
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)
 
type message struct{
	id int32
	len int32
	data []byte
}
type pack struct{
 
}
 
func(pck *pack)packHead(data []byte) *message{
	msg := &message{}
	dataio := bytes.NewReader(data)
 
 
	//：对于data这个字节数组，从低字节开始读，把读到的字节放进msg.len的地址为首的空间，因为这个空间就4个字节（int32）故而会读走4个字节
	//：dataio流的指向在读完之后，会指向data的倒数第5个字节
	binary.Read(dataio,binary.LittleEndian,&msg.id)
	binary.Read(dataio,binary.LittleEndian,&msg.len)
	//binary.Read(dataio,binary.LittleEndian,&msg.data)
	fmt.Println("----",msg)
 
	return msg
}
func(pck *pack)unpack(msg *message) []byte{
 
	databufio := bytes.NewBuffer([]byte{})
 
	binary.Write(databufio,binary.LittleEndian,msg.id)
	binary.Write(databufio,binary.LittleEndian,msg.len)
	binary.Write(databufio,binary.LittleEndian,msg.data)
	return databufio.Bytes()
}
 
 
func main(){
	bindata := []byte{}
	msg := &message{
		id:1,
		len:4,
		data:[]byte{'h','a','h','h'},
	}
	pck := &pack{}
	bindata = pck.unpack(msg)
	fmt.Println("msg - conv - bin-len:",len(bindata))
	fmt.Println(bindata)
 
 
	msg2 := pck.packHead(bindata)
	beginIndex := unsafe.Sizeof(msg2.id)+unsafe.Sizeof(msg2.len)
	msg2.data = make([]byte,4)
	binary.Read(bytes.NewReader(bindata[beginIndex:]),binary.LittleEndian,&msg2.data)
	fmt.Println("bin-conv-msg2-msg2.data=",msg2)
}
