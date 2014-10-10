package main

import "fmt"

type Human struct{
   name string
   age  int
   phone string
}

type Student struct{
   Human
   school string
}

type Employee struct{
  Human
  company string
}

func (h *Human) SayHi(){
  fmt.Printf("Hi,I am %s you can call me on %s\n",h.name,h.phone)
}

func (e *Employee) SayHi(){
  fmt.Printf("Hi,I am %s,I work at %s.Call me on %s\n",e.name,e.company,e.phone)
}

func main(){
  mark:=Student{Human{"Mark",24,"123456"},"BMD"}
  anne:=Employee{Human{"Anne",22,"1212123"},"3O"}
  mark.SayHi()
  anne.SayHi()
}
