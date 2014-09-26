package main
import "fmt"

type person struct{
  Name string
  Age int
}
func main(){
 a:=person()
 a.Name="jim"
 a.Age=19
 fmt.Println(a)
}
