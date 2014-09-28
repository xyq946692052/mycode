package main

 // 网络编程

import (
      "fmt"
      "net/http"
      )

func main(){
     http.HandleFunc("/",hello)
     http.ListenAndServe("localhost:8000",nil)
}
func hello(w http.ResponseWriter,r *http.Request){
    fmt.Fprintln(w,"hello,Gophers!Welcome to 3o!")
}
