//modify string

package main

import "fmt"

func main(){
 s:="abcdefg"
 fmt.Println("---------")
 fmt.Print("|")
 fmt.Print(s)
 fmt.Print("|")
 fmt.Println()
 c:=[]byte(s)
 c[0]='s'
 s2:=string(c)
fmt.Println("---------")
  fmt.Print("|")
 fmt.Print(s2)
 fmt.Print("|")

 fmt.Println()
fmt.Println("---------")
}
