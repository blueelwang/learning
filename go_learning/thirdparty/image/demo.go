package main

import (
	"fmt"
	"time"

	"github.com/disintegration/imaging"
)

func main() {

	t1 := time.Now()
	src, err := imaging.Open("./src_image.jpg")
	t2 := time.Now()
	fmt.Printf("Open() cost: %s\n", t2.Sub(t1))		// 612.572914ms
	if err != nil {
		fmt.Printf("Open failed, err:%s\n", err)
		return
	}

	// 宽高
	srcBounds := src.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()
	t3 := time.Now()
	fmt.Printf("src w:%d, h:%d\n", srcW, srcH)		// src w:4032, h:3024
	fmt.Printf("Bounds() cost: %s\n", t3.Sub(t2))	// 64.159µs
	
	// 缩放 & 裁切
	image := imaging.Fill(src, 1000, 1000, imaging.Center, imaging.Lanczos)
	t4 := time.Now()
	fmt.Printf("Fill() cost: %s\n", t4.Sub(t3))		// 1.148510527s

	// 保存
	err = imaging.Save(image, "./image.jpg")
	t5 := time.Now()
	fmt.Printf("%s\n",t5.Sub(t4))					// 83.543641ms
	if err != nil {
		fmt.Printf("Save failed, err:%s\n", err)
		return
	}

	// 3.98MB => 242KB
}

func demo() {
	
}