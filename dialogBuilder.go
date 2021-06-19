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
	// c.RangeListParse(root.ChildElements(), manager)
	vt1 := c.RangeListParse(root.ChildElements(), manager)
	manager.R1 = append(manager.R1, vt1...)

	// x1.Item = vt1
	// x2.Item = vt2
	// for _, v := range vt1 {
	// 	if v2, ok := v.(*XMLContainerUI); ok {
	// 		fmt.Println("v2==========------", v2, ok)
	// 		// 	for _, v3 := range v2.Item {
	// 		// 		if v4, ok := v3.(*XMLControlUI); ok {

	// 		// 			fmt.Println("XMLControlUI", v4)
	// 		// 		}
	// 		// 	}
	// 	}
	// }
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

func (c *DialogBuilder) RangeListParse(el []*etree.Element, manager *PaintManagerUI) (r1 []interface{}) {
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
			pc := NewXMLControl()

			for _, v2 := range v.Attr {
				fmt.Printf("Control :[%+v]\n", v2)
				pc.SetAttr(v2)
			}
			pc.SetPaint(manager)

			vt1 := c.RangeListParse(vElement, manager)
			pc.Item = append(pc.Item, vt1...)
			// x2.Item = vt2
			// pc.Item =
			r1 = append(r1, pc)
			// r2 = append(r2, x2)

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

			pc := NewXMLContainer()
			// var c1 *ContainerUI
			for _, v2 := range v.Attr {
				fmt.Printf("Control :[%+v]\n", v2)
				pc.SetAttr(v2)
			}
			// fmt.Printf("Container CoreUI  :[%+v]\n", pc.DoControlUI.XML)

			vt1 := c.RangeListParse(vElement, manager)
			pc.Item = append(pc.Item, vt1...)
			r1 = append(r1, pc)

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
