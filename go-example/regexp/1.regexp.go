package main

import (
    "fmt"
    "regexp"
)

func main(){
    // func MatchString(pattern string, s string) (matched bool, err error)
    // MatchString类似Match，但匹配对象是字符串。

    match, _:=regexp.MatchString("foo.*", "seafood")
    fmt.Println(match) //true

    match, _=regexp.MatchString("p([a-z]+)ch","peach")
    fmt.Println(match)//true

    // func Compile(expr string) (*Regexp, error)
    // Compile解析并返回一个正则表达式。如果成功返回，该Regexp就可用于匹配文本。

    // func (re *Regexp) MatchString(s string) bool
    // MatchString类似Match，但匹配对象是字符串。

    r, _ := regexp.Compile("p([a-z]+)ch") // x+             匹配一个或多个 x，优先匹配更多(贪婪)
    fmt.Println(r.MatchString("peach"))//true

    // func (re *Regexp) FindString(s string) string
    // Find返回保管正则表达式re在b中的最左侧的一个匹配结果的字符串。
    // 如果没有匹配到，会返回""；但如果正则表达式成功匹配了一个空字符串，也会返回""。
    // 如果需要区分这种情况，请使用FindStringIndex 或
    fmt.Println(r.FindString("peach punch")) //peach
    fmt.Println(r.FindString("pch punch")) //punch
    fmt.Println(r.FindString("pch pch")) // ""
    
    //func (re *Regexp) FindStringIndex(s string) (loc []int)
    //Find返回保管正则表达式re在b中的最左侧的一个匹配结果的起止位置的切片（显然len(loc)==2）。
    //匹配结果可以通过起止位置对b做切片操作得到：b[loc[0]:loc[1]]。如果没有匹配到，会返回nil。
    fmt.Println(r.FindStringIndex("peach punch")) //[0 5]
    fmt.Println(r.FindStringIndex("pch punch"))  //[4 9]

    // func (re *Regexp) FindStringSubmatch(s string) []string
    // Find返回一个保管正则表达式re在b中的最左侧的一个匹配结果以及（可能有的）分组匹配的结果的[]string切片。如果没有匹配到，会返回nil。
    fmt.Println(r.FindStringSubmatch("peach punch")) //[peach ea]






}