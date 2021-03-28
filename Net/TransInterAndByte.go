package Net

import (
	"bytes"
	"encoding/binary"
)

func IntToByte(n int)[]byte  {
	data:=int64(n)
	byteBuf := bytes.NewBuffer([]byte{})
	binary.Write(byteBuf,binary.BigEndian,data)
	return byteBuf.Bytes()
}

func BytesToInt(bts []byte)int {
	byteBuff := bytes.NewBuffer(bts)
	var data int64
	binary.Read(byteBuff,binary.BigEndian,&data)
	return int(data)
}