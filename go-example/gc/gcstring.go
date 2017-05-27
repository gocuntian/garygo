package main

import "runtime"
import "time"

type T struct {
	ptr **int
	pad [120]byte
}

var things []interface{}

func main() {
	setup()
	runtime.GC()
	runtime.GC()
	time.Sleep(10 * time.Millisecond)
	runtime.GC()
	runtime.GC()
	time.Sleep(10 * time.Millisecond)
}

func setup() {
	var Ts []interface{}
	buf := make([]byte, 128)

	for i := 0; i < 10000; i++ {
		s := string(buf)
		t := &T{ptr: new(*int)}
		runtime.SetFinalizer(t.ptr, func(**int) {
			panic("*int freed too early")
		})
		Ts = append(Ts, t)
		things = append(things, s[len(s):])
	}
	things = append(things, Ts...)
}

// func SetFinalizer(x, f interface{})
// SetFinalizer将x的终止器设置为f。当垃圾收集器发现一个不能接触的（即引用计数为零，程序中不能再直接或间接访问该对象）具有终止器的块时，它会清理该关联（对象到终止器）并在独立go程调用f(x)。这使x再次可以接触，但没有了绑定的终止器。如果SetFinalizer没有被再次调用，下一次垃圾收集器将视x为不可接触的，并释放x。

// SetFinalizer(x, nil)会清理任何绑定到x的终止器。

// 参数x必须是一个指向通过new申请的对象的指针，或者通过对复合字面值取址得到的指针。参数f必须是一个函数，它接受单个可以直接用x类型值赋值的参数，也可以有任意个被忽略的返回值。如果这两条任一条不被满足，本函数就会中断程序。

// 终止器会按依赖顺序执行：如果A指向B，两者都有终止器，且它们无法从其它方面接触，只有A的终止器执行；A被释放后，B的终止器就可以执行。如果一个循环结构包含一个具有终止器的块，该循环不能保证会被当垃圾收集，终止器也不能保证会执行；因为没有尊重依赖关系的顺序。

// x的终止器会在x变为不可接触之后的任意时间被调度执行。不保证终止器会在程序退出前执行，因此一般终止器只用于在长期运行的程序中释放关联到某对象的非内存资源。例如，当一个程序丢弃一个os.File对象时没有调用其Close方法，该os.File对象可以使用终止器去关闭对应的操作系统文件描述符。但依靠终止器去刷新内存中的I/O缓冲如bufio.Writer是错误的，因为缓冲不会在程序退出时被刷新。

// 如果*x的大小为0字节，不保证终止器会执行。
