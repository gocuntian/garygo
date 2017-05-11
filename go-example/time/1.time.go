package main

import (
    "fmt"
    "time"
)
// time.Sleep(time.Duration)
// 表示睡多少时间，睡觉时，是阻塞状态

// time.After(time.Duration)
// 和Sleep差不多，意思是多少时间之后，但在取出管道内容前不阻塞\

// time.AfterFunc(time.Duration,func())
// 和After差不多，意思是多少时间之后在goroutine line执行函数
func main(){
   
      funSleep()
      funAfter()
      funAfterFunc()
      funAfterFunc2()
      funTick()
      funTime()
      funFFF()

}

func funSleep(){
    fmt.Println("start sleeping...")
    time.Sleep(time.Second)
    fmt.Println("end sleep.")
    //【结果】打印start sleeping后，等了正好1秒后，打印了end sleep
    //会阻塞，Sleep时，什么事情都不会做
}

// func After(d Duration) <-chan Time
func funAfter(){
    fmt.Println("the 1")
    //返回一个time.C这个管道，1秒(time.Second)后会在此管道中放入一个时间点(time.Now())
    //时间点记录的是放入管道那一刻的时间值
    tc:=time.After(time.Second)

    fmt.Println("the 2")
    fmt.Println("the 3")
    <-tc //阻塞中，直到取出tc管道里的数据

    fmt.Println("the 4")
    //【结果】立即打印123，等了1秒不到一点点的时间，打印了4，结束
    //打印the 1后，获得了一个空管道，这个管道1秒后会有数据进来
    //打印the 2，（这里可以做更多事情）
    //打印the 3
    //等待，直到可以取出管道的数据（取出数据的时间与获得tc管道的时间正好差1秒钟）
    //打印the 4
}


func funAfterFunc(){
    f := func(){
        fmt.Println("Time out")
    }
    time.AfterFunc(1*time.Second,f)
    //要保证主线比子线“死的晚”，否则主线死了，子线也等于死了
    time.Sleep(2 * time.Second)
    //【结果】运行了1秒后，打印出timeout，又过了1秒，程序退出
    //将一个间隔和一个函数给AfterFunc后
    //间隔时间过后，执行传入的函数
}
// 由于f函数不是在Main Line执行的，而是注册在goroutine Line里执行的
// 所以一旦后悔的话，需要使用Stop命令来停止即将开始的执行，如果已经开始执行就来不及了
func funAfterFunc2(){
    is_stop := true
    f := func(){
        fmt.Println("Time out2")
    }
    ta := time.AfterFunc(2*time.Second,f)
    time.Sleep(time.Second)
    if is_stop {
        ta.Stop()
    }
    time.Sleep(3 * time.Second)//要保证主线比子线“死的晚”，否则主线死了，子线也等于死了
    //【结果】运行了3秒多一点点后，程序退出，什么都不打印
    //注册了个f函数，打算2秒后执行
    //过了1秒后，后悔了，停掉（Stop）它

}


// func Tick(d Duration) <-chan Time
//和After差不多，意思是每隔多少时间后，其他与After一致
func funTick(){
  fmt.Println("the one111")
  tc:=time.Tick(time.Second)//返回一个time.C这个管道，1秒(time.Second)后会在此管道中放入一个时间点，
                        //1秒后再放一个，一直反复，时间点记录的是放入管道那一刻的时间
  for i:=1;i<=2;i++{
      <-tc
      fmt.Println("hello")
  }
 //每隔1秒，打印一个hello                    
}

// time.Time的方法（time.Time自己独有的函数）
// Before & After方法
// 判断一个时间点是否在另一个时间点的前面（后面），返回true或false
//func (Time) Equal
// func (t Time) Equal(u Time) bool
// 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。

// func (Time) Before
// func (t Time) Before(u Time) bool
// 如果t代表的时间点在u之前，返回真；否则返回假。

// func (Time) After
// func (t Time) After(u Time) bool
func funTime(){
    t1:=time.Now()
    time.Sleep(time.Second)
    t2:=time.Now()
    a:=t2.After(t1) //true
    fmt.Println(a)
    b := t2.Before(t1) //false
    fmt.Println(b)
}


// func (t Time) Add(d Duration) Time
// Add返回时间点t+d。

// func (t Time) AddDate(years int, months int, days int) Time
// AddDate返回增加了给出的年份、月份和天数的时间点Time。例如，时间点January 1, 2011调用AddDate(-1, 2, 3)会返回March 4, 2010。

// AddDate会将结果规范化，类似Date函数的做法。因此，举个例子，给时间点October 31添加一个月，会生成时间点December 1。（从时间点November 31规范化而来）

// func (t Time) Sub(u Time) Duration
// 返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。

// func (t Time) Round(d Duration) Time
// 返回距离t最近的时间点，该时间点应该满足从Time零值到该时间点的时间段能整除d；如果有两个满足要求的时间点，距离t相同，会向上舍入；如果d <= 0，会返回t的拷贝。

// func (t Time) Truncate(d Duration) Time
// 类似Round，但是返回的是最接近但早于t的时间点；如果d <= 0，会返回t的拷贝。
func funFFF(){
    //两个时间点相减，获得时间差（Duration）
    t1 :=time.Now()
    time.Sleep(time.Second)
    t2:=time.Now()
    d:=t2.Sub(t1) ////时间2减去时间1
    fmt.Println(d)////打印结果差不多为1.000123几秒，因为Sleep无法做到精确的睡1秒 后发生的时间 
    

    //Add拿一个时间点，add一个时长，获得另一个时间点 减去   先发生时间，是正数

    t3:=time.Now() // //现在是12点整（假设）,那t1记录的就是12点整
    t4:=t3.Add(time.Hour) //  //那t1的时间点 **加上(Add)** 1个小时，是几点呢？
    fmt.Println(t4) //  //13点（呵呵）
}