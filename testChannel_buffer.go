package main

import (
  "fmt"
  "strconv"
  "time"
)

func shooting(msg_chan chan string){
  var group = 1
  for{
    for i:=1;i<=10;i++{
      msg_chan<-strconv.Itoa(group)+":"+strconv.Itoa(i)
      }
      group++
      time.Sleep(time.Second*10)
  }
}

func count(msg_chan chan string){
  for{
    fmt.Println(<-msg_chan)
  }
}

func main(){
  var c=make(chan string)
  go shooting(c)
  go count(c)
  var input string
  fmt.Scanln(&input)
}
