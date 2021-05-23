package gd

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"syscall"
	"unsafe"

	"github.com/fjs-icu/win"
)

var initUI uint32
var defaultWndProcPtr uintptr

// 统一window 操作方法.
type Window interface {
	// 返回当前的句柄
	Handle() win.HWND
	// 返回一个可以操作的实例
	AsWindowBase() *WindowBase
	Run()
	WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr
}

type WindowCfg struct {
	Window    Window
	Parent    Window
	ClassName string
	Style     uint32
	ExStyle   uint32
	Bounds    Rect
}

type WindowBase struct {
	Hwnd         win.HWND
	Window       Window
	Visible      bool // 是否隐藏
	Enabled      bool // 是否禁用
	Name         string
	PaintManager PaintManagerUI
}

func InitWindow(window, parent Window, className string, style, exStyle uint32) error {
	return initWindowWithCfg(&WindowCfg{
		Window:    window,
		Parent:    parent,
		ClassName: className,
		Style:     style,
		ExStyle:   exStyle,
	})
}

func initWindowWithCfg(cfg *WindowCfg) error {
	if atomic.CompareAndSwapUint32(&initUI, 0, 1) {
		runtime.LockOSThread()

		var initCtrls win.INITCOMMONCONTROLSEX
		initCtrls.DwSize = uint32(unsafe.Sizeof(initCtrls))
		initCtrls.DwICC = win.ICC_LINK_CLASS | win.ICC_LISTVIEW_CLASSES | win.ICC_PROGRESS_CLASS | win.ICC_TAB_CLASSES | win.ICC_TREEVIEW_CLASSES
		win.InitCommonControlsEx(&initCtrls)
		defaultWndProcPtr = syscall.NewCallback(defaultWndProc)

	}

	wb := cfg.Window.AsWindowBase()
	wb.Window = cfg.Window
	wb.Enabled = true
	wb.Visible = cfg.Style&win.WS_VISIBLE != 0

	var windowName *uint16
	if len(wb.Name) != 0 {
		windowName = syscall.StringToUTF16Ptr(wb.Name)
	}

	var hwndParent win.HWND

	if cfg.Parent != nil {
		hwndParent = cfg.Parent.Handle()

	}

	if hwnd := cfg.Window.Handle(); hwnd == 0 {
		var x, y, w, h int32
		if cfg.Bounds.IsZero() {
			x = 200
			y = 200
			w = 500
			h = 500
		} else {
			x = int32(cfg.Bounds.X)
			y = int32(cfg.Bounds.Y)
			w = int32(cfg.Bounds.W)
			h = int32(cfg.Bounds.H)
		}
		fmt.Println(cfg.ClassName, x, y, w, h)
		// 注册窗口类
		MustRegWinProcPtr(cfg.ClassName, defaultWndProcPtr)
		// cfg.Style|win.WS_CLIPSIBLINGS
		// 创建窗口
		wb.Enabled = true

		// unsafe.Pointer(wb) 创建后
		wb.Hwnd = win.CreateWindowEx(
			cfg.ExStyle,
			syscall.StringToUTF16Ptr(cfg.ClassName),
			windowName,
			cfg.Style|win.WS_CLIPSIBLINGS,
			x,
			y,
			w,
			h,
			hwndParent,
			0,
			0,
			unsafe.Pointer(wb))
		if wb.Hwnd == 0 {
			return NewErr("CreateWindowEx")
		}
		fmt.Println("wb hwnd", wb.Hwnd)

	} else {
		wb.Hwnd = hwnd
	}

	ustyle := win.GetWindowStyle(wb.Hwnd)
	fmt.Println("ustyle : ", ustyle)

	Hwnd2WindowBase[wb.Hwnd] = wb

	return nil
}

func WindowFromHandle(hwnd win.HWND) Window {
	if wb := Hwnd2WindowBase[hwnd]; wb != nil {
		return wb.Window
	}

	return nil
}

type WindowBase2 struct {
	Hwnd  win.HWND
	Index int
}

func defaultWndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) (result uintptr) {
	// defer func() {
	// }()

	var wb *WindowBase

	if msg == win.WM_NCCREATE {
		fmt.Println("msg win.WM_NCCREATE ")
		var p2 *win.CREATESTRUCT

		p2 = (*win.CREATESTRUCT)(unsafe.Pointer(lParam))
		// fmt.Println("size ", unsafe.Sizeof(*p2))
		w := (*WindowBase)(unsafe.Pointer(p2.CreateParams))
		if w == nil {
			fmt.Println("w nil aaa", unsafe.Pointer(w))
			panic("CreateParams nil")
		}
		win.SetWindowLongPtr(hwnd, win.GWLP_USERDATA, uintptr(unsafe.Pointer(w)))
		// 设置后可收到 WM_CREATE 消息.
	} else {
		w2 := win.GetWindowLongPtr(hwnd, win.GWLP_USERDATA)
		wb = (*WindowBase)(unsafe.Pointer(w2))

	}
	if wb == nil {
		fmt.Println("DefWindowProc=================================", hwnd)
		return win.DefWindowProc(hwnd, msg, wParam, lParam)
	}

	return wb.WndProc(hwnd, msg, wParam, lParam)

}

func (c *WindowBase) Handle() win.HWND {
	return c.Hwnd
}

func (c *WindowBase) AsWindowBase() *WindowBase {
	return c
}
func (c *WindowBase) Run() {
	win.ShowWindow(c.Hwnd, win.SW_SHOW)
	win.UpdateWindow(c.Hwnd)

	var msg win.MSG
	for {
		if win.GetMessage(&msg, 0, 0, 0) == win.TRUE {
			// ustyle := win.GetWindowStyle(msg.HWnd)
			// fmt.Println("ustyle : ", ustyle)
			// if !c.PaintManager.TranslateMessage(&msg) {
			// 要处理一些事情
			// }

			win.TranslateMessage(&msg)
			win.DispatchMessage(&msg)

		} else {
			break
		}
	}
	fmt.Println("WindowBase end")
}

// 事件

func (c *WindowBase) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	window := WindowFromHandle(hwnd)
	_ = window
	// asBase := window.AsWindowBase()
	asBase := c
	// fmt.Println("msg id ", msg, hwnd)
	switch msg {
	// case win.WM_ERASEBKGND:
	// 	// WM_ERASEBKGND是在当窗口背景必须被擦除时,窗口的移动，窗口的大小的改变
	// 	return 1

	// case win.WM_HSCROLL, win.WM_VSCROLL:
	// 	if window := WindowFromHandle(win.HWND(lParam)); window != nil {
	// 		// The window that sent the notification shall handle it itself.
	// 		return window.WndProc(hwnd, msg, wParam, lParam)
	// 	}

	// case win.WM_LBUTTONDOWN, win.WM_MBUTTONDOWN, win.WM_RBUTTONDOWN:
	// 移动事件
	// wb.publishMouseEvent(&wb.mouseDownPublisher, msg, wParam, lParam)

	// case win.WM_LBUTTONUP, win.WM_MBUTTONUP, win.WM_RBUTTONUP:
	// if msg == win.WM_LBUTTONUP && wb.origWndProcPtr == 0 {
	// 	// See WM_LBUTTONDOWN for why we require origWndProcPtr == 0 here.
	// 	if !win.ReleaseCapture() {
	// 		lastError("ReleaseCapture")
	// 	}
	// }
	// wb.publishMouseEvent(&wb.mouseUpPublisher, msg, wParam, lParam)

	// case win.WM_MOUSEMOVE:
	// 	// wb.publishMouseEvent(&wb.mouseMovePublisher, msg, wParam, lParam)

	// case win.WM_MOUSEWHEEL:
	// 	// wb.publishMouseWheelEvent(&wb.mouseWheelPublisher, wParam, lParam)

	// case win.WM_SETFOCUS, win.WM_KILLFOCUS:

	// case win.WM_SETCURSOR:

	// case win.WM_CONTEXTMENU:

	// case win.WM_KEYDOWN:

	// case win.WM_KEYUP:

	// case win.WM_DROPFILES:

	// case win.WM_WINDOWPOSCHANGED:

	// case win.WM_THEMECHANGED:

	case win.WM_CREATE:
		{
			fmt.Println("WM_CREATE...")
			asBase.OnCreate(hwnd, msg, wParam, lParam)
			return 0
		}
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
		break
	}
	if asBase != nil {
		v1, v2 := asBase.PaintManager.MessageHandler(msg, wParam, lParam)
		if v1 {
			return v2
		}
	}

	return win.DefWindowProc(hwnd, msg, wParam, lParam)
}

func (c *WindowBase) OnCreate(hwnd win.HWND, msg uint32, wParam, lParam uintptr) {
	fmt.Println(" OnCreate(hwnd win.HWND, msg uint32, wParam, lParam uintptr) ")
	c.PaintManager.Init(hwnd, "")

}
