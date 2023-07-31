package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	Command()
	MultiOut()
	CombinedOutput()
}

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
