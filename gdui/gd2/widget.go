package gd2

import "github.com/lxn/win"

type WidgetBase struct {
	WindowBase
}

func NewWidgetBase() (*WidgetBase, error) {
	tt := new(WidgetBase)
	// win.WS_OVERLAPPEDWINDOW,
	err := InitWindow(tt, nil, "widgetBase", win.WS_OVERLAPPEDWINDOW, 0)

	return tt, err
}

// 方法
type Widget interface {
	Window
}

// 初始化一个 widget
func InitWidget(widget Widget, parent Window, className string, style, exStyle uint32) error {
	if parent == nil {
		return NewErr("parent cannot be nil")
	}

	if err := InitWindow(widget, parent, className, style|win.WS_CHILD, exStyle); err != nil {
		return err
	}

	return nil
}
