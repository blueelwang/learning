package systemcall

import (
	"fmt"
	"os"
)


func FileDemo() {
	// 创建文件，文件存在则清空
	file, err := os.Create("/tmp/go_file_text.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	file.WriteString("Created by Go!")
	file.Close()


	// 打开文件，不存在 err 不空
	file, err = os.Open("/tmp/go_file_text.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	buff := make([]byte, 100)
	file.ReadAt(buff, 0)
	fmt.Printf("%s\n", buff)
	file.Close()


}