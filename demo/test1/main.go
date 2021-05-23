package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"

	"github.com/fjs-icu/gd"
)

// 打印堆栈信息
func PanicTrace(kb int) string {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return string(stack)
}

func WorkRun() {
	defer func() {
		if e := recover(); e != nil {
			str := PanicTrace(1024)
			fmt.Println(str)
			os.Exit(1)
		}
	}()

	runtime.LockOSThread()

	w, err := gd.NewWidgetBase()
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = w
	w.Run()
}
func main() {
	WorkRun()
}
