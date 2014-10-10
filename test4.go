package main

import "fmt"

type Skills []string

type Human struct{
  name string
  age int
  weight int
}

type Student struct{
  Human
  Skills
  int
  speciality string
}

func main(){
  jane:=Student{Human:Human{"jane",20,110},speciality:"C#"}
  fmt.Println("Her name is ",jane.name)
  jane.Skills=[]string{"python"}
  fmt.Println("Her skills are",jane.Skills)
  jane.Skills=append(jane.Skills,"java","golang")
  fmt.Println("Now skills are :",jane.Skills)
  jane.int=3
  fmt.Println(jane.int)  

}
