package main

import (
	"d2/gd2"
	"fmt"
)

func main() {
	w, err := gd2.NewWidgetBase()
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = w
	w.Run()
}
