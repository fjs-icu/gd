package gd

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"github.com/fjs-icu/win"
)

type Include struct {
	Source string `xml:"source,attr"`
}

type Font struct {
	Shared string `xml:"shared,attr"`
	Id     string `xml:"id,attr"`
}

type Default struct {
	Shared string `xml:"shared,attr"`
}

type WindowXml struct {
	XMLName     xml.Name `xml:"Window"`
	Size        string   `xml:"size,attr"` // 最初窗口大小
	SizeBox     string   `xml:"sizebox,attr"`
	Caption     string   `xml:"caption,attr"` // 标题栏的区域(鼠标可以托动的区域)
	Roundcorner string   `xml:"roundcorner,attr"`
	Mininfo     string   `xml:"mininfo,attr"` // 窗口最小Size
	Maxinfo     string   `xml:"maxinfo,attr"` // 窗口最大Size
	Showdirty   string   `xml:"showdirty,attr"`

	Include []Include `xml:"Include"`
	Font    []Font    `xml:"Font"`
	Default []Default `xml:"Default"`
}

type WindowUI struct {
	Xml      WindowXml
	InitSize Size

	MininfoUI Size
	MaxinfoUI Size

	XMLNodeItem []interface{} //xml控件集合
	NodeItem    []interface{} //绘制控件集合

	XMLControl map[string]interface{} // 所有的控件集合
}

// 从xml-->属性
func (c *WindowUI) Parse(content []byte, paint *PaintManagerUI) {
	// var result WindowXml
	err := xml.Unmarshal(content, &c.Xml)
	if err != nil {
		fmt.Println(err)
		return
	}
	tmp := c.Xml.Size
	if tmp != "" {
		ip := GetSpLitInt32(tmp, ",")
		if len(ip) > 1 {
			c.SetInitSize(ip[0], ip[1], paint.HWndPaint)
		}
	}

	tmp = c.Xml.Mininfo
	if tmp != "" {
		ip := GetSpLitInt32(tmp, ",")
		if len(ip) > 1 {
			c.MininfoUI.Cx = ip[0]
			c.MininfoUI.Cy = ip[1]

		}
	}
	tmp = c.Xml.Maxinfo
	if tmp != "" {
		ip := GetSpLitInt32(tmp, ",")
		if len(ip) > 1 {
			c.MaxinfoUI.Cx = ip[0]
			c.MaxinfoUI.Cy = ip[1]

		}
	}

	fmt.Printf("WindowXml [%+v]\n", c)

}

// 从属性值-->xml
func (c *WindowUI) UnParse() {

}
func (c *WindowUI) SetInitSize(cx, cy int32, hWnd win.HWND) {
	c.InitSize.Cx = cx
	c.InitSize.Cy = cy
	if hWnd != 0 {
		win.SetWindowPos(hWnd, 0, 0, 0, cx, cy, win.SWP_NOZORDER|win.SWP_NOMOVE|win.SWP_NOACTIVATE)
	}
}

func GetSpLitInt32(src, splic string) (out []int32) {

	ip := strings.Split(src, splic)
	if len(ip) > 0 {
		for _, v := range ip {
			iip, err := strconv.ParseInt(v, 10, 32)
			if err == nil {
				out = append(out, int32(iip))
			}
		}
	}
	return
}

func GetSpLitInt(src, splic string) (out []int) {

	ip := strings.Split(src, splic)
	if len(ip) > 0 {
		for _, v := range ip {
			iip, err := strconv.Atoi(v)
			if err == nil {
				out = append(out, iip)
			}
		}
	}
	return
}
