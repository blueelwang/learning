package systemcall

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

// FileDemo 读文件
func FileDemo() {
	// 打开文件，不存在 err 不空
	file, err := os.Open("./data2.tmp")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	buff := make([]byte, 10)
	_, err = file.ReadAt(buff, 0)	// 指定起始位置读，文件不能已经被关闭，否则会报错
	fmt.Printf("buf: %s, err: %s\n", buff, err)

	left, err := ioutil.ReadAll(file) // 读取文件全部（当前游标位置不影响）
	fmt.Printf("all: %s, err: %s\n", left, err)	

	// 重复关闭会报错
	// file.Close()
	// err = file.Close()
	// fmt.Printf("err:%s\n", err) // err:close ./data2.tmp: file already closed
}

// FileCreateDemo 创建文件示例
func FileCreateDemo() {
	// O_CREATE 指定创建文件
	// O_EXCL 指定只在不存在时创建，否则报错且不会创建
	f, err := os.OpenFile("./data.tmp", os.O_RDWR | os.O_CREATE | os.O_EXCL, 0664)
	if err != nil {
		fmt.Println(err)	// 如果存在，报错： open ./data.tmp: file exists
	} else {
		f.Close()
	}

	// 创建文件，如果存在则清空
	f, err = os.Create("./data2.tmp")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString("Created by Go -- " + time.Now().Format("2006-01-02 15:04:05"))
}

// FileMetaDemo 读取文件无信息
func FileMetaDemo() {
	fileInfo, err := os.Stat("./data2.tmp")
	if os.IsNotExist(err) {
		fmt.Printf("file not exist, err:%+v", err)
		return
	}
	fmt.Printf("FileInfo.IsDir():\t%t\n", fileInfo.IsDir())
	fmt.Printf("FileInfo.Name():\t%s\n", fileInfo.Name())
	fmt.Printf("FileInfo.Size():\t%d\n", fileInfo.Size())
	fmt.Printf("FileInfo.ModTime():\t%s\n", fileInfo.ModTime().Format("2006-01-02 15:04:05"))
	/*
	FileInfo.IsDir():	false
	FileInfo.Name():	data2.tmp
	FileInfo.Size():	36
	FileInfo.ModTime():	2021-02-22 15:02:36
	*/
}

// DirDemo 目录操作
func DirDemo() {
	dir := "./"

	// 分批读取
	f, err := os.Open(dir)
	if err != nil {
		fmt.Printf("open dir failed, err:%+v", err)
		return
	}
	defer f.Close()

	for err != io.EOF {
		var list []os.FileInfo
		list, err = f.Readdir(2)
		if err == nil {
			for _, v := range list {
				fmt.Printf("isDir:%t, file name:%s, modTime:%s\n", v.IsDir(), v.Name(), v.ModTime().Format("2006-01-02 15:04:05"))
			}
		} else if err != nil && err != io.EOF {
			fmt.Printf("open dir failed, err:%+v", err)
			return
		}
	}


	fmt.Println("---------------------------------------------")
	fileInfo, err := os.Stat(dir)
	if os.IsNotExist(err) {
		fmt.Printf("dir not exist, err:%+v", err)
		return
	}
	fmt.Printf("isDir:%v\n", fileInfo.IsDir())

	// 一次性读出所有
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("read dir failed, err:%+v", err)
		return
	}
	for _, v := range files {
		fmt.Printf("isDir:%t, file name:%s, modTime:%s\n", v.IsDir(), v.Name(), v.ModTime().Format("2006-01-02 15:04:05"))
	}

	return
}

// FileLockDemo 文件加锁
func FileLockDemo() {
	/*
	syscall.Flock() 调用对文件加锁。
	LOCK_EX	排他锁
	LOCK_SH	共享锁
	LOCK_UN	解锁
	LOCK_NB	非阻塞加锁

	一个进程加上排他锁之后，其他进程再加锁会被阻塞（默认）或 返回失败（LOCK_NB）
	同一进程内加锁可重入，但是对于指向同一个文件的不同的文件描述符，同一进程内加锁也会相同排斥
	是否可以加锁成功，判断依据为文件描述符是否相同
	*/

	f, err := os.Open("./data2.tmp")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
    defer f.Close()

	f2, err := os.Open("./data2.tmp")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
    defer f2.Close()

	// 指向同一个文件的不同文件描述符
	fmt.Println(int(f.Fd()), int(f2.Fd())) // 5 6

	// 加排他锁
	fmt.Println("-------------------- before LOCK_EX ---------------------")
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
	// err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX)   	// 一个进程内，对同一个文件描述符多次加锁，不会被阻塞
	// err = syscall.Flock(int(f2.Fd()), syscall.LOCK_EX)  	// 但是对于指向同一文件的不同的文件描述符，多次加锁，会阻塞
	fmt.Println("-------------------- after  LOCK_EX ---------------------")
	if err != nil {
		fmt.Println(err)	// 加锁失败报错：resource temporarily unavailable
		return
    }
	
	time.Sleep(5 * time.Second)

	// 解锁
	fmt.Println("-------------------- before LOCK_UN ---------------------")
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	fmt.Println("-------------------- after  LOCK_UN ---------------------")
	if err != nil {
		fmt.Println(err)
	}

	return
}