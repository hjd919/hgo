package gom

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/gogf/gf/util/gconv"
)

func BigEndian() { // 大端序
	// 二进制形式：0000 0001 0000 0010
	// 十六进制表示：0000 0000 0000 0000 0001 0002 0003 0004
	var testInt int32 = 0x01020304 // 十六进制表示
	fmt.Printf("%d use big endian: \n", testInt)

	var aa int32 = 16909060

	fmt.Println("int32 to bytes111:", gconv.Bytes(testInt))
	fmt.Println("int32 to bytes2222:", gconv.Bytes(aa))

	var testBytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt)) //大端序模式
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes) //大端序模式的字节转为int32
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func LittleEndian() { // 小端序
	//二进制形式： 0000 0000 0000 0000 0001 0002 0003 0004
	var testInt int32 = 0x01020304 // 16进制
	fmt.Printf("%d use little endian: \n", testInt)

	var testBytes []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(testInt)) //小端序模式
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes) //小端序模式的字节转换
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}
func TestBianma(t *testing.T) {

	BigEndian()
	LittleEndian()

	// var a string = "你"
	// var c uint64 = 256
	// // var e uint8 = 97

	// b := gconv.Bytes(a)
	// k := gconv.Bytes(c)
	// log.Println(string(c))
	// d := gconv.Bytes(c)
	// log.Println(k, b, d)
	return
}

func hello() {
	InfoLog("test考试222")
}
