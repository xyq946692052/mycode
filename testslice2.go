package main

import "fmt"
func main(){
  myslice1:=make([]int,5)
  myslice2:=make([]int,5,10)
  myslice:=[]int{1,2,3,4,5}
  myslice3:=append(myslice,1,2,3)
  copy(myslice3,myslice1)
  for i:=0;i<len(myslice1);i++{
    fmt.Println("myslice1[",i,"]=",myslice1[i])
  }
  fmt.Println("--------------------")
  for i:=0;i<len(myslice2);i++{
    fmt.Println("myslice2[",i,"]=",myslice2[i])
  }
  fmt.Println("--------------------")
  for i:=0;i<len(myslice3);i++{
    fmt.Println("myslice3[",i,"]=",myslice3[i])
  }
  fmt.Println("--------------------")
  fmt.Println("myslice1's long :",len(myslice1),",space is ",cap(myslice1))
  fmt.Println("myslice2's long :",len(myslice2),",space is ",cap(myslice2))
  fmt.Println("myslice3's long :",len(myslice3),",space is ",cap(myslice3))
}
