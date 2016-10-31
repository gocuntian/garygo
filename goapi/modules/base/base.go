package base
import (
	"crypto/hmac"
    "crypto/md5"
    "crypto/rand"
    "encoding/hex"
    //"fmt"
   // "math"
   // "regexp"
   // "strings"
   // "time"
   // "unicode"
	"hash"
)
func EncodeMd5(str string)string{
    m:=md5.New()
    m.Write([]byte(str))
    return hex.EncodeToString(m.Sum(nil))
}

func GetRandomString(n int,alphabets ...byte)string{
    const alphanum="0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes=make([]byte,n)
    rand.Read(bytes)
    for i,b:=range bytes{
        if len(alphabets)==0{
            bytes[i]=alphanum[b%byte(len(alphanum))]
        }else{
            bytes[i]=alphabets[b%byte(len(alphabets))]
        }
    }
    return string(bytes)
}
func PBKDF2(password,salt []byte,iter,keyLen int,h func()hash.Hash)[]byte{
    prf:=hmac.New(h,password)
    hashLen:=prf.Size()
    numBlocks:=(keyLen + hashLen - 1)/hashLen

    var buf [4]byte
    dk:=make([]byte,0,numBlocks*hashLen)
    U:=make([]byte,hashLen)
    for block:=1;block<=numBlocks;block++{
        prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)
	// U_n = PRF(password, U_(n-1))
		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
    }
    return dk[:keyLen]
}