package systemcall

import (
	"time"
	"os"
	"os/signal"
	"fmt"
	"syscall"
)

func SignalDemo() {
	
	// Go标准库 syscall 中定义了和系统信号同名的常量
	// 这些都是 syscall.Signal 类型，一个int的别名，同时是 os.Signal 接口的实现类
	var sig syscall.Signal = syscall.SIGHUP
	fmt.Println(sig)


	sigRecv := make(chan os.Signal, 1)
	// 要自定义处理的信号
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	// 注册，收到关心的信号后会放入 sigRecv 通道中
	signal.Notify(sigRecv, sigs...)
	for i := 0; i < 3; i++ {
		sig := <-sigRecv
		fmt.Printf("Received a signal: %s\n", sig)
	}
	signal.Stop(sigRecv)	// 恢复信号的默认处理
	close(sigRecv)
	time.Sleep(10 * time.Second)
	

	
}

func SendSignal(pid int) {

	proc, _ := os.FindProcess(pid)
	proc.Signal(syscall.SIGINT)
	

}