package main

import "testing"

//要测试方法前增加"Test"开头

func TestAdd(t *testing.T){
    if (Add(2,4) == 6){
        t.Log("ok pass")
    }else{
        t.Error("error")
    }
}

func BenchmarkAdd(b *testing.B){
    // 如果需要初始化，比较耗时的操作可以这样：
	// b.StopTimer()
	// .... 一堆操作
	// b.StartTimer()
    for i:=0; i < b.N; i++{
        Add(2,8)
    }
}
// Go语言通过testing包提供自动化测试功能。包内测试只要运行命令 go test，就能自动运行符合规则的测试函数。
// Go语言测试约定规则
// 1.一般测试func TestXxx(*testing.T)
// 测试行必须Test开头，Xxx为字符串，第一个X必须大写的[A-Z]的字幕
// 为了测试方法和被测试方法的可读性，一般Xxx为被测试方法的函数名。

// 2.性能测试func BenchmarkXxx(*testing.B)
// 性能测试用Benchmark标记，Xxx同上。

// 3.测试文件名约定
// go语言测试文件名约定规则是必须以_test.go结尾，放在相同包下，为了方便代码阅读，一般go源码文件加上_test
// 比如源文件my.go 那么测试文件如果交your_test.go,her_test.go,my_test.go都可以，不过最好的还是my_test.go，方便阅读

// go test 
// //或者
// go test 路径+包名 

// //测试单个文件
// go test xxx_test.go xxx.go

// //打印详细信息
// go test -v

// //统计代码覆盖率
// go test -cover