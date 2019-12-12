package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func readMessage(r io.Reader) ([]byte, []byte, error) {
	if r == nil {
		return nil, nil, fmt.Errorf("invalid stream reader!")
	}

	//log.Println("to read txt len ...")
	txt_len_buf, err := readBytes(r, 4)
	if err != nil {
		return nil, nil, err
	}

	//log.Println("to read bin len ...")
	bin_len_buf, err := readBytes(r, 4)
	if err != nil {
		return nil, nil, err
	}

	txt_len := bytes_to_uint32(txt_len_buf)
	bin_len := bytes_to_uint32(bin_len_buf)
	//log.Println("smp buf", "txt_len_buf", txt_len_buf, "bin_len_buf", bin_len_buf)
	//log.Println("smp to read len", txt_len, bin_len)

	//log.Println("to read txt ...")
	txt_buf, err := readBytes(r, int(txt_len))
	if err != nil {
		return nil, nil, err
	}

	//log.Println("to read bin ...")
	bin_buf, err := readBytes(r, int(bin_len))
	if err != nil {
		return nil, nil, err
	}

	return txt_buf, bin_buf, nil
}

func readBytes(r io.Reader, length int) ([]byte, error) {
	if length <= 0 {
		return nil, nil
	}

	buf := make([]byte, length)
	for i := 0; i < length; {
		m, err := r.Read(buf[i:])
		if err != nil {
			return nil, err
		}
		i = i + m
	}

	return buf, nil
}

func writeMessage(w io.Writer, txt []byte, bin []byte) error {
	//log.Printf("send msg: %s\n", msg)
	if w == nil {
		return fmt.Errorf("invalid stream writer!")
	}

	txt_len := (uint32)(len(txt))
	txt_len_buf := uint32_to_bytes(txt_len)
	bin_len := (uint32)(len(bin))
	bin_len_buf := uint32_to_bytes(bin_len)

	_, err := w.Write(txt_len_buf)
	if err != nil {
		return err
	}

	_, err = w.Write(bin_len_buf)
	if err != nil {
		return err
	}

	_, err = w.Write(txt)
	if err != nil {
		return err
	}

	_, err = w.Write(bin)
	if err != nil {
		return err
	}

	return nil
}

//整形转换成字节
func uint32_to_bytes(x uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func bytes_to_uint32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
