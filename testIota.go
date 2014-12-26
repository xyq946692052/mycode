package main


import "fmt"

const(
    WHITE=iota
    BLACK
    BLUE
    RED
    YELLOW
)

type Color byte 

type Box struct {
 width,height,depth float64
 color Color
}

type BoxList []Box 

func (b Box) Volumn() float64{
   return b.width*b.height*b.depth
}

func (b *Box) setColor(c Color){
  b.color=c
}

func (b1 BoxList) BiggestsColor() Color{
  v:=0.00
  k:=Color(WHITE)
  for _,b:=range b1{
    if b.Volumn()>v{
        v=b.Volumn()
        k=b.color
    }
  }
  return k
}

func (b1 BoxList) PaintItBlack(){
  for i,_:=range b1{
    b1[i].setColor(BLACK)   
  }
}
func (c Color) String() string{
  strings:=[]string{"WHITE","BLACK","BLUE","RED","YELLOW"}
  return strings[c]
}

func main(){
  boxes:=BoxList{
    Box{4,4,4,RED},
    Box{10,10,1,YELLOW},
    Box{1,1,20,BLACK},
    Box{10,10,1,BLUE},
    Box{10,30,1,WHITE},
    Box{20,20,20,YELLOW},
  }
 fmt.Printf("we have %d boxes in our set\n",len(boxes))
 fmt.Println("The volumn of the  first one is",boxes[0].Volumn,"cm^3")
 fmt.Println("the color of the last one is",boxes[len(boxes)-1].color.String())
 fmt.Println("the biggest one is",boxes.BiggestsColor().String())

 fmt.Println("let's paint them all black")
 boxes.PaintItBlack()
 fmt.Println("the color of the second one is",boxes[1].color.String())
fmt.Println("Obviously,now,the biggest one is",boxes.BiggestsColor().String())





}













