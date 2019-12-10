package systemcall

import (
	"os"
	"fmt"
	"io"
)

func PipeDemo() {

	osPipe()
	ioPipe()

}

func osPipe() {
	// os.Pipe() 函数生成的管道在底层由系统支持，匿名管道会在缓冲区被写满之后阻塞写进程。
	// 管道是单项的，虽然reader，writer都是File类，但是不能 reader.Write() 或 writer.Read()
	// 注意：命名管道是可以被多路复用的，操作系统提供的管理不提供原子操作支持，需要自己处理
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Println("Error:", err)
		return;
	}

	data := []string{"www", ".", "daemon", "coder", ".", "com", " ", "is Amazing"}
	go func(writer *os.File, data []string) {
		for i := 0; i < len(data); i++ {
			writer.Write(([]byte)(data[i]))
		}
		writer.Close()
	}(writer, data)

	for {
		buffer := make([]byte, 20)
		n, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Println("Finished")
			break;
		}
		if err != nil {
			fmt.Println("Error:", err)
			break;
		}
		fmt.Printf("Len[%d], Content[%s]\n", n, buffer)
		/*
		多次写入的，读取时被合并
		Len[20], Content[www.daemoncoder.com ]
		Len[10], Content[is Amazing]
		*/
	}
}

func ioPipe() {
	/*
	Go在标准库io包中提供了一个有原子性操作保证的管道
	io.Pipe() 生成的管道是基于内存的，不是基于文件系统，没有缓冲区
	*/
	reader, writer := io.Pipe()

	data := []string{"www", ".", "daemon", "coder", ".", "com", " ", "is Amazing"}
	go func(writer *io.PipeWriter, data []string) {
		for i := 0; i < len(data); i++ {
			writer.Write(([]byte)(data[i]))
		}
		writer.Close()
	}(writer, data)

	for {
		buffer := make([]byte, 20)
		n, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Println("Finished")
			break;
		}
		if err != nil {
			fmt.Println("Error:", err)
			break;
		}
		fmt.Printf("Len[%d], Content[%s]\n", n, buffer)
		/*
		Len[3], Content[www]
		Len[1], Content[.]
		Len[6], Content[daemon]
		Len[5], Content[coder]
		Len[1], Content[.]
		Len[3], Content[com]
		Len[1], Content[ ]
		Len[10], Content[is Amazing]
		*/
	}
}