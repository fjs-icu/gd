package gd

import (
	"fmt"

	"github.com/fjs-icu/win"
)

// func DrawRect(hdc win.HDC,rc win.RECT,size int,dwPenColor RGB,bSrtle int){

// }

// func DrawRect(hdc win.HDC, rc win.RECT, size int) {
// 	var lb win.LOGBRUSH
// 	lb.LbStyle = win.BS_SOLID
// 	lb.LbColor = 0xff3355
// 	lb.LbHatch = 0
// 	hPen := win.HGDIOBJ(win.ExtCreatePen(win.PS_SOLID, 2, &lb, 0, nil))
// 	hOldOpen := win.SelectObject(hdc, hPen)

// 	win.SelectObject(hdc, win.GetStockObject(win.HOLLOW_BRUSH))
// 	win.Rectangle_(hdc, rc.Left, rc.Top, rc.Right, rc.Bottom)
// 	win.SelectObject(hdc, hOldOpen)
// 	win.DeleteObject(hPen)

// }

func DrawRect(hdc win.HDC, rc win.RECT, size int) {
	var lb win.LOGBRUSH
	lb.LbStyle = win.BS_SOLID
	lb.LbColor = 0xff3355
	lb.LbHatch = 0
	hPen := win.HGDIOBJ(win.ExtCreatePen(win.PS_SOLID, 2, &lb, 0, nil))
	hOldOpen := win.SelectObject(hdc, hPen)

	win.SelectObject(hdc, win.GetStockObject(win.HOLLOW_BRUSH))
	win.Rectangle_(hdc, rc.Left, rc.Top, rc.Right, rc.Bottom)
	win.SelectObject(hdc, hOldOpen)
	win.DeleteObject(hPen)

}

func DrawRectRound(hdc win.HDC, rc win.RECT, size int) {
	var lb win.LOGBRUSH
	lb.LbStyle = win.BS_SOLID
	lb.LbColor = 0xff3355
	lb.LbHatch = 0
	hPen := win.HGDIOBJ(win.ExtCreatePen(win.PS_SOLID, 2, &lb, 0, nil))
	hOldOpen := win.SelectObject(hdc, hPen)

	win.SelectObject(hdc, win.GetStockObject(win.HOLLOW_BRUSH))
	// win.Rectangle_(hdc, rc.Left, rc.Top, rc.Right, rc.Bottom)
	win.RoundRect(hdc, rc.Left, rc.Top, rc.Right, rc.Bottom, 10, 10)
	win.SelectObject(hdc, hOldOpen)
	win.DeleteObject(hPen)

}
func DrawImage(hdc win.HDC) {
	var si win.GdiplusStartupInput
	si.GdiplusVersion = 1
	if status := win.GdiplusStartup(&si, nil); status != win.Ok {
		fmt.Println("err", status)
		// return nil, newError(fmt.Sprintf("GdiplusStartup failed with status '%s'", status))
	}
	defer win.GdiplusShutdown()
	grs := new(win.GpGraphics)
	i := win.GdipCreateFromHDC(hdc, &grs)
	fmt.Println(i)
	i = win.GdipGraphicsClear(grs, win.Color_Blue)
	fmt.Println(i)
	bitmap := win.NewBitmapFromFile(`D:\code\src\github.com\gd\dog.bmp`)
	fmt.Println(bitmap.GetHeight(), bitmap.GetWidth())
	win.GdipDrawImageI(grs, bitmap.Image.Get(), 10, 10)
}
func DrawColor(hdc win.HDC, rc win.RECT, argb win.ARGB) {
	var si win.GdiplusStartupInput
	si.GdiplusVersion = 1
	if status := win.GdiplusStartup(&si, nil); status != win.Ok {
		fmt.Println("err", status)
		// return nil, newError(fmt.Sprintf("GdiplusStartup failed with status '%s'", status))
	}
	defer win.GdiplusShutdown()
	grs := new(win.GpGraphics)
	win.GdipCreateFromHDC(hdc, &grs)
	win.GdipGraphicsClear(grs, argb)
	// var bi win.BITMAPINFO
	// bi.BmiHeader.BiSize = uint32(unsafe.Sizeof(bi.BmiHeader))
	// var bi win.BITMAPV5HEADER
	// bi.BiSize = uint32(unsafe.Sizeof(bi))
	// // bi.BiWidth = int32(im.Bounds().Dx())
	// bi.BiHeight = -int32(im.Bounds().Dy())
	// bi.BiPlanes = 1
	// bi.BiBitCount = 32
	// bi.BiCompression = win.BI_BITFIELDS
	// dpm := int32(math.Round(float64(dpi) * inchesPerMeter))
	// bi.BiXPelsPerMeter = dpm
	// bi.BiYPelsPerMeter = dpm
	// The following mask specification specifies a supported 32 BPP
	// alpha format for Windows XP.
	// bi.BV4RedMask = 0x00FF0000
	// bi.BV4GreenMask = 0x0000FF00
	// bi.BV4BlueMask = 0x000000FF
	// bi.BV4AlphaMask = 0xFF000000

	// defer win.ReleaseDC(0, hdc)

	// var lpBits unsafe.Pointer

	// Create the DIB section with an alpha channel.
	// hBitmap := win.CreateDIBSection(hdc, &bi.BITMAPINFOHEADER, win.DIB_RGB_COLORS, &lpBits, 0, 0)
	// switch hBitmap {
	// case 0, win.ERROR_INVALID_PARAMETER:
	// }
}
