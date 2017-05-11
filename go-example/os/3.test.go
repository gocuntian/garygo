package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
)

// 读文件
// 在 Go 语言中，文件使用指向 os.File 类型的指针来表示的，也叫做文件句柄。
// 我们在前面章节使用到过标准输入os.Stdin 和标准输出 os.Stdout，他们的类型都是 *os.File

func main(){
    inputFile,inputError := os.Open("readme.md")
    if inputError != nil {
           fmt.Printf("An error occurred on opening the inputfile\n" +
            "Does the file exist?\n" +
            "Have you got acces to it?\n")
        return // exit the function on error
    }
    defer inputFile.Close()

    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerError := inputReader.ReadString('\n')
        if readerError == io.EOF {
            return
        }
        fmt.Printf("The input was:%s",inputString)
    }
}
// 变量 inputFile 是 *os.File 类型的。该类型是一个结构，表示一个打开文件的描述符（文件句柄）。然后，使用 os包里的 Open 函数来打开一个文件。该函数的参数是文件名，类型为 string。在上面的程序中，我们以只读模式打开 input.dat 文件。

// 如果文件不存在或者程序没有足够的权限打开这个文件，Open函数会返回一个错误：inputFile, inputError = os.Open("input.dat")。如果文件打开正常，我们就使用 defer.Close() 语句确保在程序退出前关闭该文件。然后，我们使用 bufio.NewReader 来获得一个读取器变量。

// 通过使用 bufio 包提供的读取器（写入器也类似），如上面程序所示，我们可以很方便的操作相对高层的 string 对象，而避免了去操作比较底层的字节。

// 接着，我们在一个无限循环中使用 ReadString('\n') 或 ReadBytes('\n') 将文件的内容逐行（行结束符 '\n'）读取出来。

// 注意： 在之前的例子中，我们看到，Unix和Linux的行结束符是 \n，而Windows的行结束符是 \r\n。在使用 ReadString 和ReadBytes 方法的时候，我们不需要关心操作系统的类型，直接使用 \n 就可以了。另外，我们也可以使用 ReadLine()方法来实现相同的功能。

// 一旦读取到文件末尾，变量 readerError 的值将变成非空（事实上，常亮 io.EOF 的值是 true），我们就会执行 return语句从而退出循环。