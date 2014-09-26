package main
import "fmt"
func main(){
 s := "abc"
 c := []byte(s)
 for k,v:=range s{
  fmt.Println(k,v)
}
