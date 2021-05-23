package gd

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
)

// 解析xml

type DialogBuilder struct {
}

// 读取xml 解析数据
// 可以做遍历两个树,一个是整个控件的xml属性,一个是真实的值,还有一个map存储所有的控件(有一个sid:唯一的值.)
func (c *DialogBuilder) Create(sfilexml string, manager *PaintManagerUI) {
	fmt.Println("test1.xml start")
	//
	content, err := os.ReadFile(sfilexml)
	if err != nil {
		fmt.Println("test1.xml ", err)
	}
	// manager.RootXml = &result
	if manager.WindowUI == nil {
		manager.WindowUI = new(WindowUI)
	}
	manager.WindowUI.Parse(content, manager)
	// manager.RootXml = &result
	c.Parse(sfilexml, manager, &manager.WindowUI.Xml)
}

// 解析子控件数据
func (c *DialogBuilder) Parse(sfilexml string, manager *PaintManagerUI, xw *WindowXml) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(sfilexml); err != nil {
		panic(err)
	}
	root := doc.SelectElement("Window")
	fmt.Println("ROOT element:", root.Tag)
	c.RangeListParse(root.ChildElements())
	vt1, _vt2 := c.RangeListParse(root.ChildElements())
	manager.R1 = vt1
	manager.R2 = _vt2

	// x1.Item = vt1
	// x2.Item = vt2
	for _, v := range vt1 {
		if v2, ok := v.(*XMLContainerUI); ok {
			fmt.Println("v2==========------", v2, ok)
			// 	for _, v3 := range v2.Item {
			// 		if v4, ok := v3.(*XMLControlUI); ok {

			// 			fmt.Println("XMLControlUI", v4)
			// 		}
			// 	}
		}
	}
	// fmt.Println("vt2", vt2)

	// fmt.Println(xw.NodeItem)
	// for _, book := range root.SelectElements("VerticalLayout") {
	// 	var pc ContainerUI
	// 	manager.RootXml.ControlUI = append(manager.RootXml.ControlUI, pc)
	// 	fmt.Println("CHILD element:", book.Tag)
	// 	if title := book.SelectElement("Slider"); title != nil {
	// 		lang := title.SelectAttrValue("name", "unknown")
	// 		fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
	// 	}
	// 	for _, attr := range book.Attr {
	// 		fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
	// 	}
	// }
}

func (c *DialogBuilder) RangeListParse(el []*etree.Element) (r1 []interface{}, r2 []interface{}) {
	for _, v := range el {
		var vElement []*etree.Element
		for _, v2 := range v.Child {
			if v3, ok := v2.(*etree.Element); ok {
				vElement = append(vElement, v3)
			}

		}

		// fmt.Printf("[%+v\n]", v)
		switch v.Tag {
		// 遍历子节点和属性
		// case "TreeNode":
		// case "Edit":
		// case "List":
		// case "Text":
		// case "Tree":
		// case "HBox":
		// case "VBox":
		// case "IList":
		// case "Combo":
		// case "Label":
		// case "Flash":
		// case "Button":
		// case "Option":
		// case "Slider":

		case "Control":
			for _, v2 := range v.Attr {
				fmt.Printf("Control :[%+v]\n", v2)
			}
			x1 := new(XMLControlUI)
			x2 := new(ControlUI)
			r1 = append(r1, x1)
			r2 = append(r2, x2)

		// case "ActiveX":
		// case "GifAnim":
		// case "Progress":
		// case "RichEdit":
		// case "CheckBox":
		// case "ComboBox":
		// case "DateTime":
		// case "TreeView":
		// case "TreeNode":
		// case "IListItem":
		// case "Container":
		// case "TabLayout":
		// case "ScrollBar":
		case "Container":
			pc := new(XMLContainer)
			x1 := new(XMLContainerUI)
			x2 := new(ContainerUI)
			pc.ContainerUI = x2
			pc.XMLContainerUI = x1
			// var c1 *ContainerUI
			for _, v2 := range v.Attr {
				fmt.Printf("Control :[%+v]\n", v2)
				pc.SetAttr(v2)
			}
			fmt.Printf("Container CoreUI  :[%+v]\n", pc.CoreUI.XMLControlUI)

			vt1, vt2 := c.RangeListParse(vElement)
			x1.Item = append(x1.Item, vt1...)
			x2.Item = vt2
			r1 = append(r1, x1)
			r2 = append(r2, x2)

			// XMLContainer
			// fmt.Printf("Child :[%+v]\n", v.Child)

			// 生成uuid
			// v.Attr
			// pC.
			// xw.NodeItem = append(xw.NodeItem, pC)
		case "VerticalLayout":

		case "HorizontalLayout":

		}

	}
	return

}
