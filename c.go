package main

import "fmt"

func add1(a *int) int{
   *a=*a+1
   return *a
}

func main(){
  x:=1
  x2:=add1(&x)
  fmt.Println(x)
  fmt.Println(x2)
}
