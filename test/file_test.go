package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"testing"
	"unicode/utf8"
)

// 文件读写
func TestFileWR(t *testing.T) {
	fullfile := "kw"
	// read file
	// 打开文件，只读
	// 1
	/*
		file, err := os.Open(fullfile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
			// os.File.Read(), io.ReadFull() 和
			// io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
			// 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
			data, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Data as hex: %x\n", data)
			fmt.Printf("Data as string: %s\n", data)
			fmt.Println("Number of bytes read:", utf8.RuneCountInString(string(data)))
	*/
	// 2
	// 读取文件到byte slice中
	fileBytes, err := ioutil.ReadFile(fullfile)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Data read: %s\n", fileBytes)

	// find kw
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`\[(.+?)\]`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(string(fileBytes), -1)
	stopKws := []string{}
	for _, r := range result1 {
		if utf8.RuneCountInString(r[1]) >= 2 { // 长度大于等于2
			stopKws = append(stopKws, r[0]+"\n")
		}
	}

	// write stop file
	err = ioutil.WriteFile(fullfile, []byte(strings.Join(stopKws, "")), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
