package main

import "fmt"

func main(){
  myArray:=[10]int{1,3,6,8,3,2,0,6,4,12}
  myslice :=myArray[:5]
  myslice1 :=myArray[5:]
  myslice2 :=myArray[:]

  fmt.Println("myArray:")
  for _,v:=range myArray{
    fmt.Print(v," ")
  } 
  fmt.Println("-----------------")  
  fmt.Println("mySlice:")
  for _,v:=range myslice{
    fmt.Print(v," ")
  }
  fmt.Println("-----------------")
  fmt.Println("mySlice1:")
  for _,v:=range myslice1{
    fmt.Print(v," ")
  }
  fmt.Println("-----------------")
  fmt.Println("mySlice2:")
  for _,v:=range myslice2{
    fmt.Print(v," ")
  }


 fmt.Println()
}
