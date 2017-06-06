package main

import (
	"fmt"
	"unsafe"
)
//案例A：unsafe.Pointer和uintptr之间的转换
//不要出现在同一个表达式中
func illegalUseA(){
	fmt.Println("===================== illegalUseA")

	pa := new([4]int)
	//拆分合法使用
     // p1：= unsafe.Pointer（uintptr（unsafe.Pointer（pa））+ unsafe.Sizeof（pa [0]））
     //进两个表达式（非法使用）：
	ptr := uintptr(unsafe.Pointer(pa))
    p1 := unsafe.Pointer(ptr + unsafe.Sizeof(pa[0]))
	//“go vet”会对上面的一行发出警告：
     //可能滥用unsafe.Pointer

     //不安全的软件包文档，https：//golang.org/pkg/unsafe/#Pointer，
     //认为以上分裂是非法的。
     //但是当前的Go编译器和运行时（1.7.3）无法检测
     //这个非法使用。
     //但是，为了使您的程序运行良好的后期Go版本，
     //最好遵守不安全的软件包文档。
	  *(*int)(p1) = 123
    fmt.Println("*(*int)(p1)  :", *(*int)(p1))
}

// case B：指针指向未知地址
func illegalUseB() {
    fmt.Println("===================== illegalUseB")

    a := [4]int{0, 1, 2, 3}
    p := unsafe.Pointer(&a)
    p = unsafe.Pointer(uintptr(p) + uintptr(len(a)) * unsafe.Sizeof(a[0]))
   //现在p指向由值a占用的内存的末尾。
     //到目前为止，虽然p是无效的，但没有问题。
     //但是如果我们修改p指向的值是非法的
    *(*int)(p) = 123
    fmt.Println("*(*int)(p)  :", *(*int)(p)) // 123 or not 123
    //当前的Go编译器/运行时（1.7.3）和“go vet”
     //这里不会检测到非法使用。

     //但是，当前的Go运行时（1.7.3）将会
     //检测以下代码的非法使用和恐慌。
    p = unsafe.Pointer(&a)
    for i := 0; i <= len(a); i++ {
        *(*int)(p) = 123 // Go runtime (1.7.3) never panic here in the tests

        fmt.Println(i, ":", *(*int)(p))
        //上一行的恐慌进行最后的迭代，i== 4。
         //运行时错误：无效内存地址或nil指针取消引用

        p = unsafe.Pointer(uintptr(p) + unsafe.Sizeof(a[0]))
    }
}

func main() {
    illegalUseA()
    illegalUseB()
}