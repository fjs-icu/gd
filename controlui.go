package gd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	"github.com/fjs-icu/win"
)

type DoControlUI interface {
	Paint(hdc win.HDC, rcPaint win.RECT) bool
	DoPaint(hdc win.HDC, rcPaint win.RECT) bool
}

type XMLControlUI struct {
	Pos     string `xml:"pos,attr"`
	Padding string `xml:"padding,attr"`

	Name             string `xml:"name,attr"`
	Bkcolor          string `xml:"bkcolor,attr"`
	Bkcolor1         string `xml:"bkcolor1,attr"`
	Bkcolor2         string `xml:"bkcolor2,attr"`
	Bkcolor3         string `xml:"bkcolor3,attr"`
	BorderColor      string `xml:"bordercolor,attr"`
	FocusBorderColor string `xml:"focusbordercolor,attr"`
	Colorhsl         string `xml:"colorhsl,attr"`
	BorderSize       string `xml:"bordersize,attr"`
	BorderStyle      string `xml:"borderstyle,attr"`
	BorderRound      string `xml:"borderround,attr"`
	Bkimage          string `xml:"bkimage,attr"`

	Width     string `xml:"width,attr"`
	Height    string `xml:"height,attr"`
	MinWidth  string `xml:"minwidth,attr"`
	MinHeight string `xml:"minheight,attr"`
	MaxWidth  string `xml:"maxwidth,attr"`
	MaxHeight string `xml:"maxheight,attr"`
	Text      string `xml:"text,attr"`
	Tooltip   string `xml:"tooltip,attr"`

	UserData   string `xml:"userdata,attr"`
	Tag        string `xml:"tag,attr"`
	Enabled    string `xml:"enabled,attr"`
	Mouse      string `xml:"mouse,attr"`
	KeyHoard   string `xml:"keyboard,attr"`
	Visible    string `xml:"visible,attr"`
	Float      string `xml:"float,attr"`
	Menu       string `xml:"menu,attr"`
	Virtualwnd string `xml:"virtualwnd,attr"`

	Item []interface{}
}

type XMLControl struct {
	ID  string        // 序号
	XML *XMLControlUI // xml属性
	UI  *ControlUI    // 控件真实属性

}

func (c *XMLControl) SetAttr(attr etree.Attr) {
	switch attr.Key {
	case "pos":
		c.XML.Pos = attr.Value
	case "padding":
		c.XML.Padding = attr.Value
	case "bkcolor":
		c.XML.Bkcolor = attr.Value
		c.UI.Bkcolor = attr.Value

	case "bkcolor1":
		c.XML.Pos = attr.Value
	case "bkcolor2":
		c.XML.Pos = attr.Value
	case "bkcolor3":
		c.XML.Pos = attr.Value
	case "bordercolor":
		c.XML.Pos = attr.Value
	case "focusbordercolor":
		c.XML.Pos = attr.Value
	case "colorhsl":
		c.XML.Pos = attr.Value
	case "bordersize":
		c.XML.Pos = attr.Value
	case "borderstyle":
		c.XML.Pos = attr.Value
	case "borderround":
		c.XML.Pos = attr.Value
	case "bkimage":
		c.XML.Bkimage = attr.Value
	case "width":
		c.XML.Pos = attr.Value
	case "height":
		c.XML.Pos = attr.Value
	case "minwidth":
		c.XML.Pos = attr.Value
	case "minheight":
		c.XML.Pos = attr.Value
	case "maxwidth":
		c.XML.Pos = attr.Value
	case "maxheight":
		c.XML.Pos = attr.Value
	case "name":
		c.XML.Name = attr.Value
	case "text":
		c.XML.Pos = attr.Value
	case "tooltip":
		c.XML.Pos = attr.Value
	case "userdata":
		c.XML.Pos = attr.Value
	case "tag":
		c.XML.Pos = attr.Value
	case "enabled":
		c.XML.Pos = attr.Value
	case "mouse":
		c.XML.Pos = attr.Value
	case "keyboard":
		c.XML.Pos = attr.Value
	case "visible":
		c.XML.Pos = attr.Value
	case "float":
		c.XML.Pos = attr.Value
	case "menu":
		c.XML.Pos = attr.Value
	case "virtualwnd":
		c.XML.Pos = attr.Value

	}
}

type ControlUI struct {
	DoControlUI
	Cover   *ControlUI
	Bkcolor string

	Item []interface{}
}

func String2Int16(src string) uint64 {
	numberStr := strings.Replace(src, "0x", "", -1)
	numberStr = strings.Replace(src, "#", "", -1)

	numberStr = strings.Replace(numberStr, "0X", "", -1)
	n, err := strconv.ParseUint(numberStr, 16, 64)
	if err != nil {
		fmt.Println("ControlUI DoPaint", err)

		panic(err)
	}
	return n

}
func (c *ControlUI) Paint(hdc win.HDC, rcPaint win.RECT) bool {
	fmt.Println("ControlUI DoPaint")
	fmt.Println("ControlUI ============================")

	if !c.DoPaint(hdc, rcPaint) {
		return false
	}
	if c.Cover != nil {
		return c.Cover.Paint(hdc, rcPaint)
	}
	return true
}

func (c *ControlUI) DoPaint(hdc win.HDC, rcPaint win.RECT) bool {
	// 绘制 背景颜色-->背景图-->状态图-->文本-->边框
	fmt.Println("ControlUI DoPaint")
	icolor := String2Int16(c.Bkcolor)
	fmt.Println("ControlUI icolor ", icolor)

	DrawColor(hdc, rcPaint, win.ARGB(icolor))
	return true
}

// 面板绘制
type XMLContainer struct {
	ID     string          // 序号
	XML    *XMLContainerUI // xml属性
	UI     *ContainerUI    // 控件真实属性
	CoreUI *XMLControl
}

func (c *XMLContainer) SetAttr(attr etree.Attr) {
	// inset mousechild vscrollbarstyle hscrollbar hscrollbarstyle childpadding childalign  childvalign
	switch attr.Key {
	case "inset":
		{
			c.XML.Inset = attr.Value
		}
	case "mousechild":
		{
			c.XML.Inset = attr.Value
		}
	case "vscrollbarstyle":
		{
			c.XML.Inset = attr.Value
		}
	case "hscrollbar":
		{
			c.XML.Inset = attr.Value
		}
	case "hscrollbarstyle":
		{
			c.XML.Inset = attr.Value
		}
	case "childpadding":
		{
			c.XML.Inset = attr.Value
		}
	case "childalign":
		{
			c.XML.Inset = attr.Value
		}
	case "childvalign":
		{
			c.XML.Inset = attr.Value
		}
	default:
		{
			if c.CoreUI == nil {
				c.CoreUI = new(XMLControl)
				c.CoreUI.XML = &c.XML.CoreUI
				c.CoreUI.UI = &c.UI.UI

			}
			c.CoreUI.SetAttr(attr)

		}
	}

}

type XMLContainerUI struct {
	CoreUI XMLControlUI

	Inset           string `xml:"inset,attr"`
	Mousechild      string `xml:"mousechild,attr"`
	Vscrollbarstyle string `xml:"vscrollbarstyle,attr"`
	Hscrollbar      string `xml:"hscrollbar,attr"`
	Hscrollbarstyle string `xml:"hscrollbarstyle,attr"`
	Childpadding    string `xml:"childpadding,attr"`
	Childvalign     string `xml:"childvalign,attr"`
	Childalign      string `xml:"childalign,attr"`

	Item []interface{}
}

type ContainerUI struct {
	UI ControlUI
	// CoreUI ControlUI

	// 绘制子节点
	Item []interface{}
}

func (c *ContainerUI) Paint(hdc win.HDC, rcPaint win.RECT) bool {
	fmt.Println("ContainerUI DoPaint")

	if !c.DoPaint(hdc, rcPaint) {
		return false
	}
	if c.UI.Cover != nil {
		return c.UI.Paint(hdc, rcPaint)
	}
	return true
}

func (c *ContainerUI) DoPaint(hdc win.HDC, rcPaint win.RECT) bool {
	// 绘制 背景颜色-->背景图-->状态图-->文本-->边框
	fmt.Println("ContainerUI DoPaint")
	// DrawColor(hdc, rcPaint, 0xffff0000)
	c.UI.Paint(hdc, rcPaint)
	return true
}
