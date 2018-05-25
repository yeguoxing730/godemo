package file

import (
	"log"
	"fmt"
	"os"
	"io"
	"bufio"
	"io/ioutil"
)

//按字节读取，将整个文件读取到缓冲区buffer
func test1() {
	file,err := os.Open("filetoread.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	fileinfo,err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileinfo.Size()
	buffer := make([]byte,fileSize)

	bytesread,err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bytes read:",bytesread)
	fmt.Println("bytestream to string:",string(buffer))
}

//每次读取固定字节
//问题容易出现乱码，因为中文和中文符号不占一个字符
func test2() {
	const BufferSize = 100
	file,err := os.Open("filetoread.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte,BufferSize)

	for {
		bytesread,err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}

		fmt.Println("bytes read:",bytesread)
		fmt.Println("bytestream to string:",string(buffer[:bytesread]))
	}
}


//逐行读取
func test3() {
	file ,err := os.Open("filetoread.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	/*
		ScanLines (默认)
		ScanWords
		ScanRunes (遍历UTF-8字符非常有用)
		ScanBytes
	*/

	//是否有下一行
	for scanner.Scan() {
		fmt.Println("read string:",scanner.Text())
	}

}

//使用ioutil读取文件的所有内容
func test4() {
	bytes,err := ioutil.ReadFile("filetoread.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("total bytes read：",len(bytes))
	fmt.Println("string read:",string(bytes))
}

func test5()  {
	file ,err := os.Open("filetoread.txt")
	if err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}

}