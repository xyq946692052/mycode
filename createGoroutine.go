package main

import (
  "fmt"
 )

func List_elem(n int){
 for i:=0;i<n;i++{
   fmt.Println(i) 
 }
}

func main(){
 go List_elem(100)
 var input string
 fmt.Scanln(&input)
}

