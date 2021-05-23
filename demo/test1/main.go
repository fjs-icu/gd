package main

import (
	"fmt"
	"runtime"

	"github.com/fjs-icu/gd"
)

func main() {
	runtime.LockOSThread()

	w, err := gd.NewWidgetBase()
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = w
	w.Run()
}
