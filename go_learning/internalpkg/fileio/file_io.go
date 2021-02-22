package fileio

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func FileIODemo() {
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// str, err := reader.ReadString('\n') // 读到一个换行就结束
		str, isPrefix, err := reader.ReadLine()
		if err == io.EOF { // io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Printf("%s\n", str)
		fmt.Println(isPrefix)
	}

}