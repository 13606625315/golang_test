package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "time"
)

// https://github.com/golang/protobuf/blob/master/ptypes/timestamp/timestamp.pb.go
type Timestamp struct {
    // Represents seconds of UTC time since Unix epoch
    // 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
    // 9999-12-31T23:59:59Z inclusive.
    Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
    // Non-negative fractions of a second at nanosecond resolution. Negative
    // second values with fractions must still have non-negative nanos values
    // that count forward in time. Must be from 0 to 999,999,999
    // inclusive.
    Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    // XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}
type proto_head struct{
	text_len int32
	bin_len int32
}


func main() {
    buff := new(bytes.Buffer)

    ts := &Timestamp{
        Seconds: time.Now().Unix(),
        Nanos:   0,
    }
	head:=&proto_head{text_len:3, bin_len: 0}
    err := binary.Write(buff, binary.LittleEndian, ts)
    err = binary.Write(buff, binary.LittleEndian, head)
    if err != nil {
        panic(err)
    }
    fmt.Println("done=",buff.Bytes())
}