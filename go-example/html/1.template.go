package main

import (
    "html/template"
    "os"
    "io/ioutil"
    "time"
    "fmt"
)

func main(){
    t:=template.New("第一个模板").Delims("[[", "]]") //创建一个模板,设置模板边界
    t, _ = t.Parse("hello,[[.UserName]]\n")   //解析模板文件
    data :=map[string]interface{}{"UserName":template.HTML("<script>alert('you have been pwned')</script>")}
    t.Execute(os.Stdout,data)  //执行模板的merger操作，并输出到控制台

    t2 := template.New("新模板")  //创建模板
    t2.Funcs(map[string]interface{}{"tihuan":tihuan}) //向模板中注入函数
    bytes, _ := ioutil.ReadFile("test2.html")  //读文件
    template.Must(t2.Parse(string(bytes))) //将字符串读作模板
    t2.Execute(os.Stdout,map[string]interface{}{"UserName":"你好世界"})
    fmt.Println("\n",t2.Name(),"\n")

    t3, _ :=template.ParseFiles("test1.html") //将一个文件读作模板
    t3.Execute(os.Stdout,data)
    fmt.Println(t3.Name(),"\n") //模板名称

    t4, _ := template.ParseGlob("test1.html")
    t4.Execute(os.Stdout,data)
    fmt.Println(t4.Name())
}

func tihuan(str string) string{
    return str + "--------"+ time.Now().Format("2006-01-02")
}

// type Template struct {
//     // 底层的模板解析树，会更新为HTML安全的
//     Tree *parse.Tree
//     // 内含隐藏或非导出字段
// }
// func New(name string) *Template
// 创建一个名为name的模板
// func (t *Template) Delims(left, right string) *Template
// Delims方法用于设置action的分界字符串，应用于之后的Parse、ParseFiles、ParseGlob方法。
// 嵌套模板定义会继承这种分界符设置。空字符串分界符表示相应的默认分界符：{{或}}。返回值就是t，以便进行链式调用
// func (t *Template) Parse(src string) (*Template, error)
// Parse方法将字符串text解析为模板。嵌套定义的模板会关联到最顶层的t。
// Parse可以多次调用，但只有第一次调用可以包含空格、注释和模板定义之外的文本。
// 如果后面的调用在解析后仍剩余文本会引发错误、返回nil且丢弃剩余文本；
// 如果解析得到的模板已有相关联的同名模板，会覆盖掉原模板。。

// func (t *Template) Execute(wr io.Writer, data interface{}) error
// Execute方法将解析好的模板应用到data上，并将输出写入wr。
// 如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。
// 模板可以安全的并发执行。

// type HTML string
// HTML用于封装一个已知安全的HTML文档片段。它不应被第三方使用，也不能用于含有未闭合的标签或注释的HTML文本
//。该类型适用于封装一个效果良好的HTML生成器生成的HTML文本或者本包模板的输出的文本。

// func (t *Template) Funcs(funcMap FuncMap) *Template
// Funcs方法向模板t的函数字典里加入参数funcMap内的键值对。
// 如果funcMap某个键值对的值不是函数类型或者返回值不符合要求会panic。
// 但是，可以对t函数列表的成员进行重写。方法返回t以便进行链式调用。

// func ParseFiles(filenames ...string) (*Template, error)
// ParseFiles函数创建一个模板并解析filenames指定的文件里的模板定义。
// 返回的模板的名字是第一个文件的文件名（不含扩展名），内容为解析后的第一个文件的内容。
// 至少要提供一个文件。如果发生错误，会停止解析并返回nil。

// func ParseGlob(pattern string) (*Template, error)
// ParseGlob创建一个模板并解析匹配pattern的文件（参见glob规则）里的模板定义。
// 返回的模板的名字是第一个匹配的文件的文件名（不含扩展名），内容为解析后的第一个文件的内容。
// 至少要存在一个匹配的文件。如果发生错误，会停止解析并返回nil。
// ParseGlob等价于使用匹配pattern的文件的列表为参数调用ParseFiles。

