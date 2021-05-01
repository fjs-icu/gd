package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

func main() {
	hInst := win.GetModuleHandle(nil)
	hIcon := win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_APPLICATION))
	hCursor := win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW))

	var wc win.WNDCLASSEX
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	wc.LpfnWndProc = syscall.NewCallback(wndProc)
	wc.HInstance = hInst
	wc.HIcon = hIcon
	wc.HCursor = hCursor
	wc.HbrBackground = win.COLOR_WINDOW + 1
	wc.LpszClassName = syscall.StringToUTF16Ptr("go windwow")
	wc.Style = win.CS_HREDRAW | win.CS_VREDRAW
	win.RegisterClassEx(&wc)

	hWnd := win.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr("go windwow"),
		syscall.StringToUTF16Ptr("go windwow"),
		win.WS_OVERLAPPEDWINDOW,
		400,
		200,
		640,
		480,
		0,
		0,
		hInst,
		nil)
	win.ShowWindow(hWnd, win.SW_SHOW)
	win.UpdateWindow(hWnd)

	var msg win.MSG
	for {
		if win.GetMessage(&msg, 0, 0, 0) == win.TRUE {
			win.TranslateMessage(&msg)
			win.DispatchMessage(&msg)
		} else {
			break
		}
	}
}

func wndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) (result uintptr) {

	var ps win.PAINTSTRUCT
	switch msg {
	case win.WM_LBUTTONDOWN:
		// MsgBox("鼠标右键点击", "Win32_Mouse", win.MB_OK)

		break
	// case WM_RBUTTONDOWN:
	// break

	case win.WM_PAINT:
		hdc := win.BeginPaint(hwnd, &ps)

		var lb win.LOGBRUSH
		lb.LbStyle = win.BS_SOLID
		lb.LbColor = 0xff000
		lb.LbHatch = 0

		hPen := win.HGDIOBJ(win.ExtCreatePen(win.PS_SOLID, 2, &lb, 0, nil))
		hOldOpen := win.SelectObject(hdc, hPen)

		var pt win.POINT
		win.MoveToEx(hdc, 0, 0, &pt)
		win.LineTo(hdc, 100, 100)
		win.EndPaint(hwnd, &ps)

		win.SelectObject(hdc, hOldOpen)
		win.DeleteObject(hPen)
		fmt.Println("paint....")

		break
	case win.WM_CREATE:
		{
			fmt.Println("create")
			OnCreate(hwnd, msg, wParam, lParam)
			brak
		}

	case win.WM_DESTROY:
		win.PostQuitMessage(0)
		break
	default:
		return win.DefWindowProc(hwnd, msg, wParam, lParam)
	}
	return 0
}

type PaintManagerUI struct {
	HWndPaint     win.HWND
	Name          string
	HDCPaint      win.HDC
	HDcOffscreen  win.HDC
	HDcBackground win.HDC
}

func (c *PaintManagerUI) Init(hWnd win.HWND, pstrName string) {
	// 移除之前所有的控件

	//初始化
	c.Name = pstrName
	// if( m_hWndPaint != hWnd ) {
	// 	m_hWndPaint = hWnd;
	// 	m_hDcPaint = ::GetDC(hWnd);
	// 	m_aPreMessages.Add(this);
	// }
	if c.HWndPaint != hWnd {
		c.HWndPaint = hWnd
		c.HDCPaint = win.GetDC(hWnd)
	}

}
