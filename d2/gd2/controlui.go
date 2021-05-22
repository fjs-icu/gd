package gd2

import (
	"fmt"

	"github.com/beevik/etree"
	"github.com/fjs-icu/win"
)

type DoControlUI interface {
	Paint(hdc win.HDC, rcPaint win.RECT) bool
	DoPaint(hdc win.HDC, rcPaint win.RECT) bool
}

type XMLControl struct {
	ID           string        // 序号
	XMLControlUI *XMLControlUI // xml属性
	ControlUI    *ControlUI    // 控件真实属性
}

func (c *XMLControl) SetAttr(attr etree.Attr) {
	switch attr.Key {
	case "pos":
		c.XMLControlUI.Pos = attr.Value
	case "padding":
		c.XMLControlUI.Pos = attr.Value
	case "bkcolor":
		c.XMLControlUI.Pos = attr.Value
	case "bkcolor1":
		c.XMLControlUI.Pos = attr.Value
	case "bkcolor2":
		c.XMLControlUI.Pos = attr.Value
	case "bkcolor3":
		c.XMLControlUI.Pos = attr.Value
	case "bordercolor":
		c.XMLControlUI.Pos = attr.Value
	case "focusbordercolor":
		c.XMLControlUI.Pos = attr.Value
	case "colorhsl":
		c.XMLControlUI.Pos = attr.Value
	case "bordersize":
		c.XMLControlUI.Pos = attr.Value
	case "borderstyle":
		c.XMLControlUI.Pos = attr.Value
	case "borderround":
		c.XMLControlUI.Pos = attr.Value
	case "bkimage":
		c.XMLControlUI.Pos = attr.Value
	case "width":
		c.XMLControlUI.Pos = attr.Value
	case "height":
		c.XMLControlUI.Pos = attr.Value
	case "minwidth":
		c.XMLControlUI.Pos = attr.Value
	case "minheight":
		c.XMLControlUI.Pos = attr.Value
	case "maxwidth":
		c.XMLControlUI.Pos = attr.Value
	case "maxheight":
		c.XMLControlUI.Pos = attr.Value
	case "name":
		c.XMLControlUI.Name = attr.Value
	case "text":
		c.XMLControlUI.Pos = attr.Value
	case "tooltip":
		c.XMLControlUI.Pos = attr.Value
	case "userdata":
		c.XMLControlUI.Pos = attr.Value
	case "tag":
		c.XMLControlUI.Pos = attr.Value
	case "enabled":
		c.XMLControlUI.Pos = attr.Value
	case "mouse":
		c.XMLControlUI.Pos = attr.Value
	case "keyboard":
		c.XMLControlUI.Pos = attr.Value
	case "visible":
		c.XMLControlUI.Pos = attr.Value
	case "float":
		c.XMLControlUI.Pos = attr.Value
	case "menu":
		c.XMLControlUI.Pos = attr.Value
	case "virtualwnd":
		c.XMLControlUI.Pos = attr.Value

	}
}

type XMLControlUI struct {
	Pos  string `xml:"pos,attr"`
	Name string `xml:"name,attr"`

	Item []interface{}
}

type ControlUI struct {
	DoControlUI
	Cover *ControlUI
	Item  []interface{}
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

	return true
}

// 面板绘制
type XMLContainer struct {
	ID             string          // 序号
	XMLContainerUI *XMLContainerUI // xml属性
	ContainerUI    *ContainerUI    // 控件真实属性
	CoreUI         *XMLControl
}

func (c *XMLContainer) SetAttr(attr etree.Attr) {
	// inset mousechild vscrollbarstyle hscrollbar hscrollbarstyle childpadding childalign  childvalign
	switch attr.Key {
	case "inset":
		{
			c.XMLContainerUI.Inset = attr.Value
		}
	case "mousechild":
		{
			c.XMLContainerUI.Inset = attr.Value
		}
	case "vscrollbarstyle":
		{
			c.XMLContainerUI.Inset = attr.Value
		}
	case "hscrollbar":
		{
			c.XMLContainerUI.Inset = attr.Value
		}
	case "hscrollbarstyle":
		{
			c.XMLContainerUI.Inset = attr.Value
		}
	case "childpadding":
		{
			c.XMLContainerUI.Inset = attr.Value
		}
	case "childalign":
		{
			c.XMLContainerUI.Inset = attr.Value
		}
	case "childvalign":
		{
			c.XMLContainerUI.Inset = attr.Value
		}
	default:
		{
			if c.CoreUI == nil {
				c.CoreUI = new(XMLControl)
				c.CoreUI.XMLControlUI = &c.XMLContainerUI.CoreUI
				c.CoreUI.ControlUI = &c.ContainerUI.CoreUI

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
	CoreUI ControlUI
	// 绘制子节点
	Item []interface{}
}

func (c *ContainerUI) Paint(hdc win.HDC, rcPaint win.RECT) bool {
	fmt.Println("ContainerUI DoPaint")

	if !c.DoPaint(hdc, rcPaint) {
		return false
	}
	if c.CoreUI.Cover != nil {
		return c.CoreUI.Paint(hdc, rcPaint)
	}
	return true
}

func (c *ContainerUI) DoPaint(hdc win.HDC, rcPaint win.RECT) bool {
	// 绘制 背景颜色-->背景图-->状态图-->文本-->边框
	fmt.Println("ContainerUI DoPaint")
	DrawColor(hdc, rcPaint, 0xffff0000)
	return true
}
