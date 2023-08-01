package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// https://darjun.github.io/2022/11/01/godailylib/osexec/
// 执行外部命令的固定流程：
// 1 调用exec.Command()创建命令对象；
// 2 调用Cmd.Run()执行命令
// Tips：中途可以尝试利用 cmd.Stdin 获取输入; cmd.Stdout/cmd.Stderr获取输出和错误

func main() {}

// Command 执行命令输出结果到控制台
func Command() {
	cmd := exec.Command("cal")
	// Output out and err to io.Writer
	cmd.Stdout = os.Stdout // out of cmd -> os
	cmd.Stderr = os.Stderr // err of cmd -> os
	// RUn the command
	if err := cmd.Run(); err != nil {
		log.Fatalf("exec cal failed: %v\n", err)
	}
}

// MultiOut 输出到多个目的地
func MultiOut() {
	// io.writers
	f, _ := os.OpenFile("out.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	buf := bytes.NewBuffer(nil)
	w := io.MultiWriter(f, buf) // integrate f & buf into MultiWriter

	cmd := exec.Command("cal")
	cmd.Stdout = w
	cmd.Stderr = w
	if err := cmd.Run(); err != nil {
		log.Fatalf("exec cal failed: %v\n", err)
	}

	fmt.Println(buf.String())
}

// CombinedOutput 直接运行命令并获得运行结果
func CombinedOutput() {
	cmd := exec.Command("cal")
	output, _ := cmd.CombinedOutput() // 运行命令并获得输出结果
	fmt.Println(string(output))
}

// Stdin cmd.Stdin 从 io.Reader 读取输入
func Stdin() {
	cmd := exec.Command("cat")
	cmd.Stdin = bytes.NewBufferString("hello world!")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("exec cal failed: %v\n", err)

	}
}

// Env cmd.Env 将环境变量写入cmd
func Env() {
	cmd := exec.Command("echo", "$NAME", "$AGE")
	cmd.Env = append(os.Environ(), "NAME=test", "AGE=18")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}
	fmt.Println(string(out))
}

// Path LookPath 查找命令的路径，如果路径存在则代表命令存在
func Path() {
	path, err := exec.LookPath("ls")
	if err != nil {
		log.Fatalf("look path of ls failed: %v\n", err)
	}
	log.Printf("path of ls is %v\n", path)

	path, err = exec.LookPath("bad-cmd")
	if err != nil {
		log.Fatalf("look path of bad-cmd failed: %v\n", err)
	}
	log.Printf("path of bad-cmd is %v\n", path)
}
