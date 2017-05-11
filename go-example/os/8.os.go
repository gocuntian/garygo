package main

import (
    "log"
    "os"
)
//Check Read and Write Permissions(检查文件读写权限)：
// func IsPermission(err error) bool
// 返回一个布尔值说明该错误是否表示因权限不足要求被拒绝。ErrPermission和一些系统调用错误会使它返回真。
func main(){
    file, err := os.OpenFile("text.txt",os.O_WRONLY, 0666)
    if err != nil {
        if os.IsPermission(err) {
            log.Println("Error: Write permission denied.")
        }
    }
    file.Close()

    file, err = os.OpenFile("text.txt",os.O_RDONLY,0666)
    if err != nil {
        if os.IsPermission(err){
            log.Println("Error: read permission denied.")
        }
    }
    file.Close()
}