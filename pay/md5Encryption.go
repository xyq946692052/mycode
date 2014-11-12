package main
/*
 *md5加密
 */
import (
  "crypto/md5"
  "encoding/hex"
  "fmt"
 )

func main(){
  h:=md5.New()
  h.Write([]byte("123456"))
  str:=hex.EncodeToString(h.Sum(nil))
  fmt.Printf("%s\n",str)
}
