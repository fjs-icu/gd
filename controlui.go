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
}

type XMLControl struct {
	DoControlUI
	ID       string        // 序号
	XML      *XMLControlUI // xml属性
	UI       *ControlUI    // 控件真实属性
	HDCPaint *PaintManagerUI
	Item     []interface{}
}

func NewXMLControl() *XMLControl {
	pc := new(XMLControl)
	x1 := new(XMLControlUI)
	x2 := new(ControlUI)
	pc.UI = x2
	pc.XML = x1
	return pc
}

func (c *XMLControl) Paint(hdc win.HDC, rcPaint win.RECT) bool {
	fmt.Println("ControlUI DoPaint")
	fmt.Println("ControlUI ============================")

	if !c.DoPaint(hdc, rcPaint) {
		return false
	}
	for _, v := range c.Item {
		if v2, ok := v.(DoControlUI); ok {
			v2.Paint(hdc, rcPaint)
		}
	}

	return true
}

func (c *XMLControl) DoPaint(hdc win.HDC, rcPaint win.RECT) bool {
	// 绘制 背景颜色-->背景图-->状态图-->文本-->边框
	fmt.Printf("ControlUI DoPaint [%v]\n", c.UI.Bkcolor)

	DrawColor(hdc, rcPaint, win.ARGB(c.UI.Bkcolor))
	return true
}

func (c *XMLControl) SetPaint(pa *PaintManagerUI) {
	c.HDCPaint = pa
}
func (c *XMLControl) SetAttr(attr etree.Attr) {
	// 设置属性
	fmt.Println("XMLControl SetAttr")
	va := attr.Value
	switch attr.Key {
	case "pos":
		c.XML.Pos = va
	case "padding":
		c.XML.Padding = va
	case "bkcolor":
		str := va
		str = strings.Replace(str, "0x", "", -1)
		str = strings.Replace(str, "#", "", -1)

		str = strings.Replace(str, "0X", "", -1)

		c.XML.Bkcolor = str
		icolor := String2Int16(str)

		c.UI.Bkcolor = icolor
		fmt.Println("ControlUI icolor ", icolor)

		fmt.Println("XMLControl bkcolor", str)
	case "bkcolor1":
		c.XML.Pos = va
	case "bkcolor2":
		c.XML.Pos = va
	case "bkcolor3":
		c.XML.Pos = va
	case "bordercolor":
		c.XML.Pos = va
	case "focusbordercolor":
		c.XML.Pos = va
	case "colorhsl":
		c.XML.Pos = va
	case "bordersize":
		c.XML.Pos = va
	case "borderstyle":
		c.XML.Pos = va
	case "borderround":
		c.XML.Pos = va
	case "bkimage":
		c.XML.Bkimage = va
	case "width":
		c.XML.Width = va
	case "height":
		c.XML.Pos = va
	case "minwidth":
		c.XML.Pos = va
	case "minheight":
		c.XML.Pos = va
	case "maxwidth":
		c.XML.Pos = va
	case "maxheight":
		c.XML.Pos = va
	case "name":
		c.XML.Name = va
	case "text":
		c.XML.Pos = va
	case "tooltip":
		c.XML.Pos = va
	case "userdata":
		c.XML.Pos = va
	case "tag":
		c.XML.Pos = va
	case "enabled":
		c.XML.Pos = va
	case "mouse":
		c.XML.Pos = va
	case "keyboard":
		c.XML.Pos = va
	case "visible":
		c.XML.Pos = va
	case "float":
		c.XML.Pos = va
	case "menu":
		c.XML.Pos = va
	case "virtualwnd":
		c.XML.Pos = va

	}
}

type ControlUI struct {
	Cover *ControlUI
	// Bkcolor string
	Bkcolor uint64
	XYFixed Size // width="100" + height="30"（用处：尺寸,与pos相冲突，谁在后，以谁为准。）
	Item    []interface{}
}

func String2Int16(src string) uint64 {
	if src == "" {
		return 0
	}
	n, err := strconv.ParseUint(src, 16, 64)
	if err != nil {
		fmt.Println("ParseUint DoPaint", err)

		// panic(err)
		return 0
	}
	return n

}

// 面板绘制
type XMLContainer struct {
	DoControlUI

	ID     string          // 序号
	XML    *XMLContainerUI // xml属性
	UI     *ContainerUI    // 控件真实属性
	CoreUI *XMLControl
	Item   []interface{}
}

func NewXMLContainer() *XMLContainer {
	pc := new(XMLContainer)
	x1 := new(XMLContainerUI)
	x2 := new(ContainerUI)
	x3 := NewXMLControl()
	pc.XML = x1
	pc.UI = x2
	pc.CoreUI = x3

	return pc
}

func (c *XMLContainer) Paint(hdc win.HDC, rcPaint win.RECT) bool {
	fmt.Println("ContainerUI DoPaint")

	if !c.DoPaint(hdc, rcPaint) {
		return false
	}

	for _, v := range c.Item {
		if v2, ok := v.(DoControlUI); ok {
			v2.Paint(hdc, rcPaint)
		}
	}

	return true
}

func (c *XMLContainer) DoPaint(hdc win.HDC, rcPaint win.RECT) bool {
	// 绘制 背景颜色-->背景图-->状态图-->文本-->边框
	fmt.Println("ContainerUI DoPaint START")
	// DrawColor(hdc, rcPaint, 0xffff0000)
	// c.UI.UI.Paint(hdc, rcPaint)

	c.CoreUI.Paint(hdc, rcPaint)

	fmt.Println("ContainerUI DoPaint END")

	return true
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
				c.CoreUI = NewXMLControl()
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
}

type ContainerUI struct {
	UI ControlUI
	// CoreUI ControlUI

	// 绘制子节点
	Item []interface{}
}
