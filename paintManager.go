package gd

import (
	"fmt"
	"unsafe"

	"github.com/fjs-icu/win"
)

type PaintManagerUI struct {
	HWndPaint     win.HWND
	Name          string
	HDCPaint      win.HDC
	HDcOffscreen  win.HDC
	HDcBackground win.HDC

	// Root    DoControlUI
	R1 []interface{}
	R2 []interface{}

	// RootXml *WindowXml
	WindowUI *WindowUI
	DialogBuilder
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
	// c.Root = new(ControlUI)
	build := new(DialogBuilder)
	build.Create("test1.xml", c)
}

func (c *PaintManagerUI) MessageHandler(msg uint32, wParam, lParam uintptr) (bool, uintptr) {
	if c.HWndPaint == 0 {
		return false, 0

	}
	//fmt.Println("HWndPaint ====", c.HWndPaint)

	// 处理分发事件...
	switch msg {
	case win.WM_NCHITTEST:
		{
			// 移动无标题栏的窗口(鼠标在客户区拖动移动窗口,非标题栏区域)
			return c.OnNcHitTest(msg, wParam, lParam)
		}
	// case win.WM_ERASEBKGND:
	case win.WM_PAINT:
		var ps win.PAINTSTRUCT
		hdc := win.BeginPaint(c.HWndPaint, &ps)
		defer win.EndPaint(c.HWndPaint, &ps)

		// if c.Root == nil {
		// 	// 如果没有根目录,则绘制黑色区域

		// 	DrawColor(hdc, ps.RcPaint, 0xffff0000)
		// 	return true
		// }

		// 分层
		// 刷新子区域
		// 绘制所有区域

		iSaveDc := win.SaveDC(hdc)
		// c.Root.Paint(hdc, ps.RcPaint)
		for _, v := range c.R2 {
			if v2, ok := v.(DoControlUI); ok {
				fmt.Println("-==================---")
				v2.Paint(hdc, ps.RcPaint)
				// for _, v3 := range v2.Item {
				// 	if v4, ok := v3.(*XMLControlUI); ok {
				// 		v3.
				// 			fmt.Println("XMLControlUI", v4)
				// 	}
				// }
			}
		}
		win.RestoreDC(hdc, iSaveDc)
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
		return true, 0

	case win.WM_GETMINMAXINFO:
		// 鼠标移动窗口,改变窗口大小,都会调用

		return c.OnGetMinMaxInfo(msg, wParam, lParam)

	case win.WM_SIZE:
		return c.OnSize(msg, wParam, lParam)

	}

	return false, 0
}

func (c *PaintManagerUI) OnSize(msg uint32, wParam, lParam uintptr) (bool, uintptr) {
	szRoundcorner := c.WindowUI.Roundcorner
	if !win.IsIconic(c.HWndPaint) && (szRoundcorner.Cx > 0 || szRoundcorner.Cy > 0) {
		var rcClient win.RECT
		win.GetClientRect(c.HWndPaint, &rcClient)

		hrgn := win.CreateRoundRectRgn(rcClient.Left, rcClient.Top, rcClient.Right, rcClient.Bottom, szRoundcorner.Cx, szRoundcorner.Cy)
		win.SetWindowRgn(c.HWndPaint, hrgn, true)
		win.DeleteObject(win.HGDIOBJ(hrgn))
	}
	return true, 0

}

func (c *PaintManagerUI) OnGetMinMaxInfo(msg uint32, wParam, lParam uintptr) (bool, uintptr) {
	var lpmmi *win.MINMAXINFO
	lpmmi = (*win.MINMAXINFO)(unsafe.Pointer(lParam))
	if lpmmi != nil {
		if c.WindowUI.MininfoUI.Cx > 0 {
			lpmmi.PtMinTrackSize.X = c.WindowUI.MininfoUI.Cx
		}
		if c.WindowUI.MininfoUI.Cy > 0 {
			lpmmi.PtMinTrackSize.Y = c.WindowUI.MininfoUI.Cy
		}
		if c.WindowUI.MaxinfoUI.Cx > 0 {
			lpmmi.PtMaxTrackSize.X = c.WindowUI.MaxinfoUI.Cx
		}
		if c.WindowUI.MaxinfoUI.Cy > 0 {
			lpmmi.PtMaxTrackSize.Y = c.WindowUI.MaxinfoUI.Cy
		}
	}
	return true, 0
}

func (c *PaintManagerUI) OnNcHitTest(msg uint32, wParam, lParam uintptr) (bool, uintptr) {
	x := int32(win.GET_X_LPARAM(lParam))
	y := int32(win.GET_Y_LPARAM(lParam))
	var pt win.POINT
	pt.X = x
	pt.Y = y
	win.ScreenToClient(c.HWndPaint, &pt)
	var rcClient win.RECT
	win.GetClientRect(c.HWndPaint, &rcClient)

	if !win.IsZoomed(c.HWndPaint) {
		rcSizeBox := c.WindowUI.SizeBox
		if pt.Y < rcClient.Top+rcSizeBox.Top {
			if pt.X < rcClient.Left+rcSizeBox.Left {
				return true, win.HTTOPLEFT
			}
			if pt.X > rcClient.Right-rcSizeBox.Right {
				return true, win.HTTOPRIGHT
			}
			return true, win.HTTOP
		} else if pt.Y > rcClient.Bottom-rcSizeBox.Bottom {
			if pt.X < rcClient.Left+rcSizeBox.Left {
				return true, win.HTBOTTOMLEFT
			}
			if pt.X > rcClient.Right-rcSizeBox.Right {
				return true, win.HTBOTTOMRIGHT
			}
			return true, win.HTBOTTOM
		}
		if pt.X < rcClient.Left+rcSizeBox.Left {
			return true, win.HTLEFT
		}
		if pt.X > rcClient.Right-rcSizeBox.Right {
			return true, win.HTRIGHT
		}
	}
	rcCaption := c.WindowUI.Caption
	if (pt.X >= rcClient.Left+rcCaption.Left) && (pt.X < rcClient.Right-rcCaption.Right) && pt.Y >= rcCaption.Top && pt.Y < rcCaption.Bottom {
		// 判断当前点是什么控件,如果是按钮等不能实现拖动控件
		// 判断当前点所在的最外层控件是否隐藏,判断当前点处于最里层的控件是什么. 暂未想好如何实现.
		return true, win.HTCAPTION

	}

	return true, win.HTCLIENT
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

func (c *PaintManagerUI) SetInitSize(cx, cy int32) {
	c.WindowUI.SetInitSize(cx, cy, c.HWndPaint)

}
