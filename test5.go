package main

import "fmt"

type Human struct{
  name string
  age int
  phone string
}

type Employee struct{
  Human
  speciality string
 phone string
}

func main(){
 Bob:=Employee{Human:Human{"Bob",12,"123456"},speciality:"C++",phone:"111111"}
 fmt.Println("Bob's work phone:",Bob.phone)
 fmt.Println("Bob's person phone:",Bob.Human.phone)
}
