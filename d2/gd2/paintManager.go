package gd2

import (
	"fmt"

	"github.com/fjs-icu/win"
)

type PaintManagerUI struct {
	HWndPaint     win.HWND
	Name          string
	HDCPaint      win.HDC
	HDcOffscreen  win.HDC
	HDcBackground win.HDC
}

func (c *PaintManagerUI) Init(hWnd win.HWND, pstrName string) {
	// 移除之前所有的控件
	fmt.Println("1====", hWnd)
	if c == nil {
		fmt.Println(" c nil")

		return
	}

	//初始化
	c.Name = pstrName
	// if( m_hWndPaint != hWnd ) {
	// 	m_hWndPaint = hWnd;
	// 	m_hDcPaint = ::GetDC(hWnd);
	// 	m_aPreMessages.Add(this);
	// }
	fmt.Println("2====")

	if c.HWndPaint != hWnd {
		fmt.Println("HWndPaint ====", hWnd)

		c.HWndPaint = hWnd
		c.HDCPaint = win.GetDC(hWnd)
	}

}

func (c *PaintManagerUI) MessageHandler(msg uint32, wParam, lParam uintptr) bool {
	if c.HWndPaint == 0 {
		return false
	}
	//fmt.Println("HWndPaint ====", c.HWndPaint)

	// 处理分发事件...
	switch msg {
	// case win.WM_ERASEBKGND:
	case win.WM_PAINT:
		var ps win.PAINTSTRUCT
		hdc := win.BeginPaint(c.HWndPaint, &ps)
		// image.draw(hdc, location)
		var rc win.RECT
		rc.Left = 10
		rc.Top = 10

		rc.Right = 200
		rc.Bottom = 200
		// size := 1
		// DrawRect(hdc, rc, size)
		// DrawRectRound(hdc, rc, size)
		DrawColor(hdc)

		// win.SetBkColor(hdc, win.COLORREF(0xff0000))
		win.EndPaint(c.HWndPaint, &ps)
		fmt.Println("MessageHandler222 paint....")
		// var ps win.PAINTSTRUCT
		// hdc := win.BeginPaint(c.HWndPaint, &ps)
		// var cl Color
		// cl = 0xFFFFFFFF
		// cl = RGB(255, 44, 0)
		// sysColor := win.COLOR_BTNFACE
		// var lb win.LOGBRUSH
		// lb.LbStyle = win.BS_SOLID
		// lb.LbColor = 0xff3300
		// lb.LbHatch = 0
		// hPen := win.HGDIOBJ(win.ExtCreatePen(win.PS_SOLID, 2, &lb, 0, nil))
		// hOldOpen := win.SelectObject(hdc, hPen)
		// win.SetBkMode(hdc, win.TRANSPARENT)

		// var bgRC win.RECT
		// win.GetWindowRect(c.HWndPaint, &bgRC)

		// var rc win.RECT
		// win.GetWindowRect(c.HWndPaint, &rc)

		// win.SetBrushOrgEx(hdc, bgRC.Left-rc.Left, bgRC.Top-rc.Top, nil)
		// win.SetBkColor(hdc, win.COLORREF(win.GetSysColor(sysColor)))
		// cc := win.GetBkColor(hdc)
		// fmt.Println("cc ", cc)
		// cc2 := win.GetBkColor(hdc)
		// fmt.Println("cc ", cc2)
		// cleardevice()
		// var ps win.PAINTSTRUCT

		// hdc := win.BeginPaint(c.HWndPaint, &ps)

		// var pt win.POINT
		// win.MoveToEx(hdc, 0, 0, &pt)
		// win.LineTo(hdc, 100, 100)
		// win.Rectangle_(hdc, 100, 100, 100, 100)

		// win.EndPaint(c.HWndPaint, &ps)

		// win.SelectObject(hdc, hOldOpen)
		// win.DeleteObject(hPen)
		// fmt.Println("paint....")
		// win.InvalidateRect(c.HWndPaint, nil, false)
		return true
		break
	}
	return false
}

var ii int

func (c *PaintManagerUI) TranslateMessage(msg *win.MSG) bool {
	// 往几个子模块发送消息
	// msg.HWnd
	// win.GetWindowSty
	fmt.Println("TranslateMessage ...")

	ustyle := win.GetWindowStyle(msg.HWnd)
	fmt.Println("ustyle : ", ustyle)
	ii++
	if ii > 2 {
		return true
	}
	return false

}
