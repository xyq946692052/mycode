package main
 
import (
   "fmt"
   "errors"
)

func main(){
  err:=errors.New("emit macho dwarf:elf header corrupted")
  if err !=nil{
     fmt.Println(err)
  }
}

