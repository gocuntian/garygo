一.指针类型：

*类型：普通指针，用于传递对象地址，不能进行指针运算。

unsafe.Pointer：通用指针类型，用于转换不同类型的指针，不能进行指针运算。

uintptr：用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。

　　unsafe.Pointer 可以和 普通指针 进行相互转换。
　　unsafe.Pointer 可以和 uintptr 进行相互转换。

　　也就是说 unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。

type ArbitraryType int
ArbitraryType在本文档里表示任意一种类型，但并非一个实际存在与unsafe包的类型。

type Pointer *ArbitraryType
Pointer类型用于表示任意类型的指针。有4个特殊的只能用于Pointer类型的操作：

1) 任意类型的指针可以转换为一个Pointer类型值
2) 一个Pointer类型值可以转换为任意类型的指针
3) 一个uintptr类型值可以转换为一个Pointer类型值
4) 一个Pointer类型值可以转换为一个uintptr类型值
因此，Pointer类型允许程序绕过类型系统读写任意内存。使用它时必须谨慎。

func Sizeof(v ArbitraryType) uintptr
Sizeof返回类型v本身数据所占用的字节数。返回值是“顶层”的数据占有的字节数。例如，若v是一个切片，它会返回该切片描述符的大小，而非该切片底层引用的内存的大小。

func Alignof(v ArbitraryType) uintptr
Alignof返回类型v的对齐方式（即类型v在内存中占用的字节数）；若是结构体类型的字段的形式，它会返回字段f在该结构体中的对齐方式。

func Offsetof(v ArbitraryType) uintptr
Offsetof返回类型v所代表的结构体字段在结构体中的偏移量，它必须为结构体类型的字段的形式。换句话说，它返回该结构起始处与该字段起始处之间的字节数。


Golang的unsafe包是一个很特殊的包。 为什么这样说呢？ 本文将详细解释。


来自go语言官方文档的警告


unsafe包的文档是这么说的：


导入unsafe的软件包可能不可移植，并且不受Go 1兼容性指南的保护。

Go 1 兼容性指南这么说：


导入unsafe软件包可能取决于Go实现的内部属性。 我们保留对可能导致程序崩溃的实现进行更改的权利。

当然包名称暗示unsafe包是不安全的。 但这个包有多危险呢？ 让我们先看看unsafe包的作用。


Unsafe包的作用


直到现在（Go1.7），unsafe包含以下资源：




三个函数：



func Alignof（variable ArbitraryType）uintptr

func Offsetof（selector ArbitraryType）uintptr

func Sizeof（variable ArbitraryType）uintptr



和一种类型：



类型Pointer * ArbitraryType


这里，ArbitraryType不是一个真正的类型，它只是一个占位符。


与Golang中的大多数函数不同，上述三个函数的调用将始终在编译时求值，而不是运行时。 这意味着它们的返回结果可以分配给常量。


（BTW，unsafe包中的函数中非唯一调用将在编译时求值。当传递给len和cap的参数是一个数组值时，内置函数和cap函数的调用也可以在编译时被求值。）


除了这三个函数和一个类型外，指针在unsafe包也为编译器服务。


出于安全原因，Golang不允许以下之间的直接转换：




两个不同指针类型的值，例如 int64和 float64。



指针类型和uintptr的值。

但是借助unsafe.Pointer，我们可以打破Go类型和内存安全性，并使上面的转换成为可能。这怎么可能发生？让我们阅读unsafe包文档中列出的规则：



任何类型的指针值都可以转换为unsafe.Pointer。

unsafe.Pointer可以转换为任何类型的指针值。

uintptr可以转换为unsafe.Pointer。

unsafe.Pointer可以转换为uintptr。

这些规则与Go规范一致：


底层类型uintptr的任何指针或值都可以转换为指针类型，反之亦然。

规则表明unsafe.Pointer类似于c语言中的void 。当然，void 在C语言里是危险的！


在上述规则下，对于两种不同类型T1和T2，可以使 T1值与unsafe.Pointer值一致，然后将unsafe.Pointer值转换为 T2值（或uintptr值）。通过这种方式可以绕过Go类型系统和内存安全性。当然，滥用这种方式是很危险的。

再来一点 unsafe.Pointer 和 uintptr


这里有一些关于unsafe.Pointer和uintptr的事实：



uintptr是一个整数类型。

即使uintptr变量仍然有效，由uintptr变量表示的地址处的数据也可能被GC回收。


unsafe.Pointer是一个指针类型。

但是unsafe.Pointer值不能被取消引用。

如果unsafe.Pointer变量仍然有效，则由unsafe.Pointer变量表示的地址处的数据不会被GC回收。




unsafe.Pointer是一个通用的指针类型，就像* int等。


由于uintptr是一个整数类型，uintptr值可以进行算术运算。 所以通过使用uintptr和unsafe.Pointer，我们可以绕过限制，* T值不能在Golang中计算偏移量：

\
unsafe包有多危险


关于unsafe包，Ian，Go团队的核心成员之一，已经确认：




在unsafe包中的函数的签名将不会在以后的Go版本中更改，



并且unsafe.Pointer类型将在以后的Go版本中始终存在。

所以，unsafe包中的三个函数看起来不危险。 go team leader甚至想把它们放在别的地方。 unsafe包中这几个函数唯一不安全的是它们调用结果可能在后来的版本中返回不同的值。 很难说这种不安全是一种危险。


看起来所有的unsafe包的危险都与使用unsafe.Pointer有关。 unsafe包docs列出了一些使用unsafe.Pointer合法或非法的情况。 这里只列出部分非法使用案例：


编译器很难检测Go程序中非法的unsafe.Pointer使用。 运行“go vet”可以帮助找到一些潜在的错误，但不是所有的都能找到。 同样是Go运行时，也不能检测所有的非法使用。 非法unsafe.Pointer使用可能会使程序崩溃或表现得怪异（有时是正常的，有时是异常的）。 这就是为什么使用不安全的包是危险的。


转换T1 为 T2


对于将 T1转换为unsafe.Pointer，然后转换为 T2，unsafe包docs说：


如果T2比T1大，并且两者共享等效内存布局，则该转换允许将一种类型的数据重新解释为另一类型的数据。

这种“等效内存布局”的定义是有一些模糊的。 看起来go团队故意如此。 这使得使用unsafe包更危险。


由于Go团队不愿意在这里做出准确的定义，本文也不尝试这样做。 这里，列出了已确认的合法用例的一小部分，


合法用例1：在[]T和[]MyT之间转换


在这个例子里，我们用int作为T：


type MyInt int

在Golang中，[] int和[] MyInt是两种不同的类型，它们的底层类型是自身。 因此，[] int的值不能转换为[] MyInt，反之亦然。 但是在unsafe.Pointer的帮助下，转换是可能的：

合法用例2: 调用sync/atomic包中指针相关的函数


sync / atomic包中的以下函数的大多数参数和结果类型都是unsafe.Pointer或*unsafe.Pointer：



func CompareAndSwapPointer（addr * unsafe.Pointer，old，new unsafe.Pointer）（swapped bool）

func LoadPointer（addr * unsafe.Pointer）（val unsafe.Pointer）

func StorePointer（addr * unsafe.Pointer，val unsafe.Pointer）

func SwapPointer（addr * unsafe.Pointer，new unsafe.Pointer）（old unsafe.Pointer）

要使用这些功能，必须导入unsafe包。
注意： unsafe.Pointer是一般类型，因此 unsafe.Pointer的值可以转换为unsafe.Pointer，反之亦然。



unsafe包用于Go编译器，而不是Go运行时。

使用unsafe作为程序包名称只是让你在使用此包是更加小心。

使用unsafe.Pointer并不总是一个坏主意，有时我们必须使用它。

Golang的类型系统是为了安全和效率而设计的。 但是在Go类型系统中，安全性比效率更重要。 通常Go是高效的，但有时安全真的会导致Go程序效率低下。 unsafe包用于有经验的程序员通过安全地绕过Go类型系统的安全性来消除这些低效。

unsafe包可能被滥用并且是危险的。