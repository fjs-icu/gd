package main

import (
	"fmt"
	"runtime"

	"github.com/jthmath/winapi"
	"github.com/lxn/win"
)

//    全局变量
var hBitmap winapi.HBITMAP

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	inst, err := winapi.GetModuleHandle("")
	if err != nil {
		fmt.Println("获取实例失败")
		return
	}

	r := WinMain(inst, "", 0)
	fmt.Println("WinMain函数返回", r)
}

func WinMain(Inst winapi.HINSTANCE, Cmd string, nCmdShow int32) int32 {
	var err error

	//    1.    注册窗口类
	atom, err := MyRegisterClass(Inst)
	if err != nil {
		fmt.Println("注册窗口类失败")
		return 0
	}
	fmt.Println("注册窗口类成功", atom)

	//    2.    创建窗口
	wnd, err := winapi.CreateWindow("主窗口类", "golang    windows    编程",
		winapi.WS_OVERLAPPEDWINDOW, 0,
		winapi.CW_USEDEFAULT, winapi.CW_USEDEFAULT, winapi.CW_USEDEFAULT, winapi.CW_USEDEFAULT,
		0, 0, Inst, 0)
	if err != nil {
		fmt.Println("创建窗口失败")
		return 0
	}
	fmt.Println("创建窗口成功", wnd)
	winapi.ShowWindow(wnd, winapi.SW_SHOW)
	winapi.UpdateWindow(wnd)

	//    3.    主消息循环
	var msg winapi.MSG
	msg.Message = winapi.WM_QUIT + 1 //    让它不等于    winapi.WM_QUIT

	for winapi.GetMessage(&msg, 0, 0, 0) > 0 {
		winapi.TranslateMessage(&msg)
		winapi.DispatchMessage(&msg)
	}

	return int32(msg.WParam)
}

func WndProc(hWnd winapi.HWND, message uint32, wParam uintptr, lParam uintptr) uintptr {
	var hTemp winapi.HANDLE

	switch message {
	case winapi.WM_CREATE:
		hTemp, _ = winapi.LoadImageByName(0, "D:\\0011.bmp",
			winapi.IMAGE_BITMAP, 0, 0, winapi.LR_LOADFROMFILE)
		hBitmap = winapi.HBITMAP(hTemp)
	case winapi.WM_PAINT:
		OnPaint(hWnd)
	case winapi.WM_DESTROY:
		winapi.PostQuitMessage(0)
	case winapi.WM_COMMAND:
		OnCommand(hWnd, wParam, lParam)
	default:
		return winapi.DefWindowProc(hWnd, message, wParam, lParam)
	}
	return 0
}

func OnPaint(hWnd winapi.HWND) {
	var err error
	// var ps winapi.PAINTSTRUCT
	var ps win.PAINTSTRUCT
	var h win.HWND
	h = (win.HWND)(hWnd)
	hdc := win.BeginPaint(h, &ps)
	// hdc, err := winapi.BeginPaint(hWnd, &ps)
	// if err != nil {
	// 	return
	// }
	// defer winapi.EndPaint(hWnd, &ps) //    defer    终于有用武之地了
	// _ = hdc
	// SetBkColor并不是用来改变背景颜色，而是设置文字输出的背景色
	win.SetBkColor(hdc, win.COLORREF(0xff0000))
	//    HDC    mdc    =    CreateCompatibleDC(hdc);
	// mdc, err := winapi.CreateCompatibleDC(hdc)
	// if err != nil {
	// 	return
	// }
	// defer winapi.DeleteDC(mdc)

	// winapi.SelectObject(mdc, winapi.HGDIOBJ(hBitmap))

	// //    这个函数的第4、5个参数分别是图片的宽、高
	// //    为了简便起见，我直接写在了这里
	// //    实际项目中当然要用过程序获取一下
	// winapi.BitBlt(hdc, 0, 0, 480, 640, mdc, 0, 0, winapi.SRCCOPY)
}

func OnCommand(hWnd winapi.HWND, wParam uintptr, lParam uintptr) {
	//    暂时不需要特殊处理    WM_COMMAND
	winapi.DefWindowProc(hWnd, winapi.WM_COMMAND, wParam, lParam)
}

func MyRegisterClass(hInstance winapi.HINSTANCE) (atom uint16, err error) {
	var wc winapi.WNDCLASS
	wc.Style = winapi.CS_HREDRAW | winapi.CS_VREDRAW
	wc.PfnWndProc = WndProc
	wc.CbClsExtra = 0
	wc.CbWndExtra = 0
	wc.HInstance = hInstance
	wc.HIcon = 0
	wc.HCursor, err = winapi.LoadCursorById(0, winapi.IDC_ARROW)
	if err != nil {
		return
	}
	wc.HbrBackground = winapi.COLOR_WINDOW + 1
	wc.Menu = uint16(0)
	wc.PszClassName = "主窗口类"
	wc.HIconSmall = 0

	return winapi.RegisterClass(&wc)
}
