package gd

import (
	"syscall"
	"unsafe"

	"github.com/fjs-icu/win"
)

// 必须先注册窗口类
var (
	RegWinClass       = make(map[string]bool)
	DefaultWndProcPtr uintptr
	Hwnd2WindowBase   = make(map[win.HWND]*WindowBase)
)

func MustRegWinProcPtr(className string, wndProcPtr uintptr) {
	MustRegWin(className, wndProcPtr, 0)
}

func MustRegWin(className string, wndProcPtr uintptr, style uint32) error {
	if RegWinClass[className] {
		return nil
	}
	hInst := win.GetModuleHandle(nil)
	if hInst == 0 {
		return NewErr("GetModuleHandle nil")
	}

	hIcon := win.LoadIcon(hInst, win.MAKEINTRESOURCE(7)) // rsrc uses 7 for app icon
	if hIcon == 0 {
		hIcon = win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_APPLICATION))
	}
	if hIcon == 0 {
		return NewErr("GetModuleHandle LoadIcon")

	}

	hCursor := win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW))
	if hCursor == 0 {
		return NewErr("GetModuleHandle LoadCursor")

	}

	var wc win.WNDCLASSEX
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	wc.LpfnWndProc = wndProcPtr
	wc.HInstance = hInst
	wc.HIcon = hIcon
	wc.HCursor = hCursor
	wc.HbrBackground = win.COLOR_BTNFACE + 1
	wc.LpszClassName = syscall.StringToUTF16Ptr(className)
	wc.Style = style

	if atom := win.RegisterClassEx(&wc); atom == 0 {
		return NewErr("GetModuleHandle RegisterClassEx")
	}

	RegWinClass[className] = true
	return nil
}
