package main
import "fmt"
 
func max(a,b int) int{
 if a>b {
   return a
 }
  return b
}

func sumAndproduct(a,b int)(sum,product int){
  sum=a+b
  product=a*b
  fmt.Println("a+b=",sum)
  fmt.Println("a*b=",product)
  return
}


func main(){
  x:=1
  y:=2
  z:=3
  fmt.Println(max(x,z))
  fmt.Println(max(y,z))
  sumAndproduct(3,5)
 
}
