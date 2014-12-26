package main

import "fmt"

type PersonInfo struct{
  ID string
  Name string
  Address string
}

func main(){
      var personDB map[string] PersonInfo
      personDB=make(map[string] PersonInfo)
      personDB["1234"]=PersonInfo{"1234","Peter","CN"}
      personDB["1"]=PersonInfo{"1","Anne","UK"}
      delete(personDB,"1")
      person,ok:=personDB["1"]
      if ok{
        fmt.Println("Found person",person.Name,"with ID 1234")
      }else {
        fmt.Println("Did not find person with ID 1234")
      }
}
