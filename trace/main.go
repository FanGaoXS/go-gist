package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	// 1 创建trace信息的输出文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	// 2 启动
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 需要调试的业务代码
	fmt.Println("Hello, world!")

	// 3 停止
	trace.Stop()
}
