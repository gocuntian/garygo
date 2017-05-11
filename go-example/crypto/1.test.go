package main

import (
    "fmt"
    "log"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "io/ioutil"
)
//文件散列校验

func main(){
    data, err:= ioutil.ReadFile("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Md5: %x\n\n",md5.Sum(data))
    fmt.Printf("Sha1: %x\n\n",sha1.Sum(data))
    fmt.Printf("sha256: %x\n\n",sha256.Sum256(data))
    fmt.Printf("sha512: %x\n\n",sha512.Sum512(data))
}

// Md5: 66613d903d08fd4ec0e1f0b75bba6720

// Sha1: 63323b8fbd20bf23b513de0a58f67a2a2a5d043c

// sha256: d268c64082a50729bcf256a502e77f30e96d7d20a4d216ef51e509449636bfc8

// sha512: 9bf9deaa2cf81fed5f65f0ebd6bdfb4035bcdf99bfe0bb247d964865f87cdcee46b798c9449484cba85487833eafc740585cbbf2e01ba9f9725be35c95d57c46
