package main

import (
    "os"
    "bufio"
    "fmt"
    "log"
)
//Read with a Scanner
func main(){
    file,err :=os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    //func NewScanner(r io.Reader) *Scanner
    //NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines。
    scanner :=bufio.NewScanner(file)
    //func (s *Scanner) Split(split SplitFunc)
    //Split设置该Scanner的分割函数。本方法必须在Scan之前调用。
    //func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
    //ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将空白（参见unicode.IsSpace）分隔的片段（去掉前后空白后）作为一个token返回。本函数永远不会返回空字符串。
    scanner.Split(bufio.ScanWords)
    //func (s *Scanner) Scan() bool
    //Scan方法获取当前位置的token（该token可以通过Bytes或Text方法获得），并让Scanner的扫描位置移动到下一个token。
    //当扫描因为抵达输入流结尾或者遇到错误而停止时，本方法会返回false。在Scan方法返回false后，Err方法将返回扫描时遇到的任何错误；除非是io.EOF，此时Err会返回nil。
    success :=scanner.Scan()
    if success == false {
        err = scanner.Err()
        if err == nil {
            log.Println("Scan completed and reached EOF")
        }else{
            log.Fatal(err)
        }
    }
    //func (s *Scanner) Text() string
    // Bytes方法返回最近一次Scan调用生成的token，会申请创建一个字符串保存token并返回该字符串。
    fmt.Println("First words found:",scanner.Text())


}