package main

import (
	"fmt"
)

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("data = ", data, ", len = ", len(data), ", cap = ", cap(data)) //data =  [0 1 2 3 4 5 6 7 8 9 10] , len =  11 , cap =  11
	v := data[6:8]                                                             //常规slice , data[6:8]，从第6位到第8位（返回6， 7），长度len为2， 最大可扩充长度cap为5（6-10）
	fmt.Println("v = ", v, ", len =", len(v), " cap = ", cap(v))               //v =  [6 7] , len = 2  cap =  5
	z := data[:6:8]                                                            //data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8
	fmt.Println("z = ", z, " len = ", len(z), ", cap = ", cap(z))              //z =  [0 1 2 3 4 5]  len =  6 , cap =  8
}
