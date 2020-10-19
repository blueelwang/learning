package systemcall

import (
	"os/exec"
	"fmt"
	"io"
	"bytes"
)

func CommandDemo() {
	cmd := exec.Command("echo", "-n", "Msg from Go")

	// 绑定命令的输出到 stdout，需要放在 cmd.Start() 之前
	stdout, err := cmd.StdoutPipe();
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 读取命令输出
	/*
	output := make([]byte, 30)
	n, err := stdout.Read(output)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Len[%d] Content[%s]\n", n, output[:n])
	*/

	// 缓冲输出
	var outputBuf bytes.Buffer
	for {
		tempOutput := make([]byte, 5)
		n, err := stdout.Read(tempOutput)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Error:", err)
				return
			}
		}
		if n > 0 {
			fmt.Printf("Receiving: %s\n", tempOutput)
			outputBuf.Write(tempOutput[:n])
		}
	}
	fmt.Printf("%s\n", outputBuf.String())



}