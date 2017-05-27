package main

import "runtime"

func mk2(){
	b := new([10000]byte)
	_ = b
}

func mk1(){mk2()}

func main(){
	for i := 0; i < 10; i++ {
		mk1()
		runtime.GC()
	}
}
//func GC() GC执行一次垃圾回收。