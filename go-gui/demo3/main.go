package main

import (
	"fmt"
	"log"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func base(root *sciter.Element) {
	//通过ID选择元素
	ul, _ := root.SelectById("list")
	//获取元素的文本
	text, _ := ul.Text()
	fmt.Println("text:", text)
	//获取元素的html
	//参数为真，则返回元素外部html
	//参数为假，则返回元素内部html
	text, _ = ul.Html(false)
	fmt.Println("html:", text)
	//获取子元素个数
	n, _ := ul.ChildrenCount()
	fmt.Println(n)
}

//动态的添加元素

func addElement(root *sciter.Element) {
	//创建一个元素
	add, _ := sciter.CreateElement("li", "555")
	//设置元素的属性
	add.SetAttr("data", "add")

	//通过标签和ID来选择元素，类似jquery
	ul, _ := root.SelectFirst("ul#list")
	err := ul.Insert(add, 3)
	if err != nil {
		log.Fatal("error:", err)
	}
	add2, _ := sciter.CreateElement("li", "")
	err2 := ul.Insert(add2, 4)
	//注意这里，元素先insert后再去设置html才有效
	//设置添加元素的html
	add2.SetHtml("<a href='http://www.baidu.com'>baidu</a>", sciter.SIH_REPLACE_CONTENT)
	if err2 != nil {
		log.Fatal("添加元素失败")
	}
}

//删除元素
func delElement(root *sciter.Element) {
	ul, _ := root.SelectFirst("ul#list")
	//获取第一个子元素，下标从0开始
	li, _ := ul.NthChild(0)
	//删除元素
	li.Delete()
	//我们也可以用css选择器直接选择要删除的元素
	//注意css里面的nth-child(n)下标从1开始
	li2, _ := root.SelectFirst("ul#list>li:nth-child(2)")
	//删除元素
	li2.Delete()
}

//修改元素
func updElement(root *sciter.Element) {
	li, _ := root.SelectFirst("ul#list>li:nth-child(1)")
	//给元素设置样式
	li.SetStyle("color", "#f00")
	//给元素设置html
	//参数一：html内容
	//参数二：html放在哪里，SIH_REPLACE_CONTENT表示替换旧内容
	li.SetHtml("<a href='http://www.baidu.com'>baidu.com</a>", sciter.SIH_REPLACE_CONTENT)
	//在最后面追加内容
	li.SetHtml("haha", sciter.SIH_APPEND_AFTER_LAST)
	//设置元素属性
	li.SetAttr("test", "test")
	li2, _ := root.SelectFirst("ul#list>li:nth-child(2)")
	//设置文本
	li2.SetText("我改我改")
}

func main() {
	//创建一个新窗口
	//这里参数一和参数二都使用的默认值
	//DefaultWindowCreateFlag = SW_TITLEBAR | SW_RESIZEABLE | SW_CONTROLS | SW_MAIN | SW_ENABLE_DEBUG
	//DefaultRect = &Rect{0, 0, 300, 400}
	w, err := window.New(sciter.DefaultWindowCreateFlag, sciter.DefaultRect)
	if err != nil {
		log.Fatal(err)
	}
	// 设置标题
	w.SetTitle("this is html")
	html := `<html>
	        <body>
	          <ul id="list" name="list">
			        <li>111</li>
                    <li>222</li>
                    <li>333</li>
			  </ul>
	         </body>
		   </html>`
	//加载html，从一个字符串变量中
	w.LoadHtml(html, "")
	//窗口获取根元素，这里应该是html
	root, _ := w.GetRootElement()
	base(root)
	addElement(root)
	delElement(root)
	updElement(root)
	//显示窗口
	w.Show()
	//运行窗口，进入消息循环
	w.Run()
}
