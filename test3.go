package main

import "fmt"

type Human struct {
    name string
    age int
    weight int
}

type Student struct{
    Human
    speciality string
}

func main(){
    mark:=Student{Human{"Mark",24,120},"Golang"}
    
    fmt.Println("His name is ",mark.name)
    fmt.Println("His age is ",mark.age)
    fmt.Println("His weight is ",mark.weight)
    fmt.Println("His speciality is ",mark.speciality)

    mark.speciality="java"
    fmt.Println("after change the speciality is ",mark.speciality)
    mark.Human=Human{"Peter",25,122}
    mark.Human.age-=1;
    fmt.Println(mark.Human.age)   
}
