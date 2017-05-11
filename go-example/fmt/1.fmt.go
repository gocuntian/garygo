package main

import (
    "fmt"
)
// Scan、Scanf和Scanln从标准输入os.Stdin读取文本；Fscan、Fscanf、Fscanln从指定的io.Reader接口读取文本；Sscan、Sscanf、Sscanln从一个参数字符串读取文本。

// Scanln、Fscanln、Sscanln会在读取到换行时停止，并要求一次提供一行所有条目；Scanf、Fscanf、Sscanf只有在格式化文本末端有换行时会读取到换行为止；其他函数会将换行视为空白。

// Scanf、Fscanf、Sscanf会根据格式字符串解析参数，类似Printf。例如%x会读取一个十六进制的整数，%v会按对应值的默认格式读取。
var (
    firstName, lastName, s string
    i int
    f float32
    intput = "56.12 / 5212 / Go"
    format = "%f / %d / %s"
)
// Hi xing cuntian !
// From the string we read:  56.12 5212 Go
func main(){
    fmt.Println("Please enter your full name: ")
    //Scanln类似Scan，但会在换行时才停止扫描。最后一个条目后必须有换行或者到达结束位置。
    fmt.Scanln(&firstName,&lastName)
    //fmt.Scanf("%s %s", &firstName, &lastName)
    fmt.Printf("Hi %s %s !\n",firstName,lastName)
    //Sscanf从字符串str扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递给本函数的参数。返回成功扫描的条目个数和遇到的任何错误。
    fmt.Sscanf(intput, format, &f,&i,&s)
    fmt.Println("From the string we read: ", f, i, s)
}
// Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
// Scanf 与其类似，除了Scanf 的第一个参数用作格式字符串，用来决定如何读取。
// Sscan 和以 Sscan 开头的函数则是从字符串读取，除此之外，与 Scanf 相同。如果这些函数读取到的结果与您预想的不同，您可以检查成功读入数据的个数和返回的错误。