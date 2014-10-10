//自定义栈类型

package main

import (
    "fmt"
    "stack" 
)

func main(){
 var s stack.Stack
 s.Push("a")
 s.Push("b")
 s.Push("c")
 s.Push(1)
 for{
   item,err:=s.Pop()
   if err!=nil{
     break
   }
fmt.Println(item)
 }

}
