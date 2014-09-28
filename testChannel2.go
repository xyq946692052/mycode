package main

import (
  "fmt"
  "time"
)

func fixed_shooting(msg_chan chan string){
  for{
    msg_chan <-"fixed shooting"
    time.Sleep(time.Second*1)
  } 
}

func three_point_shooting(msg_chan chan string){
  for{
    msg_chan <-"three point shooting"
    time.Sleep(time.Second*1)
  }
}

func main(){
  c_fixed:=make(chan string)
  c_3_point:=make(chan string)
  
  go fixed_shooting(c_fixed)
  go three_point_shooting(c_3_point)

  go func(){
    for{
       select {
          case msg1:=<-c_fixed:
             fmt.Println(msg1)
          case msg2:=<-c_3_point:
             fmt.Println(msg2)
       }
    } 
  }()
  var input string
  fmt.Scanln(&input)
}
