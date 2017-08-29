package main

import "golang.org/x/net/context"

func main() {
	ctx := context.Background()
	doSomething(ctx)
}

func doSomething(ctx context.Context) {

}

//go tool fix -diff -force=context gofix.go
// diff gofix.go fixed/gofix.go
// --- /tmp/go-fix271786484        2017-06-30 09:52:19.281620539 +0800
// +++ /tmp/go-fix124496579        2017-06-30 09:52:19.281620539 +0800
// @@ -1,6 +1,6 @@
//  package main

// -import "golang.org/x/net/context"
// +import "context"

//  func main() {
//         ctx := context.Background()

// go tool fix

// 命令Go fix会把指定代码包的所有go语言源码文件中的旧版本代码修正为新版本的代码。这里所说的版本即Go语言的版本。代码包的所有Go语言源码文件不包括其子代码包（如果有的话）中的文件。
// 修正操作包括把对旧程序调用的代码更换为对新程序调用的代码、把旧的语法更换为新的语法，等等。

// | 标记名称 | 标记描述 |
// | -diff | 不将修正后的内容写入文件，而只打印修正前后的内容的对比信息到标准输出。 |
// | -r | 只对目标源码文件做有限的修正操作。该标记的值即为允许的修正操作的名称。多个名称之间用英文半角逗号分隔。 |
// | -force | 使用此标记后，即使源码文件中的代码已经与Go语言的最新版本相匹配了，也会强行执行指定的修正操作。该标记的值就是需要强行执行的修正操作的名称，多个名称之间用英文半角逗号分隔。 |

// 在默认情况下，go tool fix命令程序会在目标源码文件上执行所有的修正操作。
