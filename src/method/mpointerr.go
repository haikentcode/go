package main 

 import "fmt"
 import "math"


type point struct {
 	x, y float64
 }

 
 func (p * point) Scale(f float64){

 	 p.x=p.x*f;
 	 p.y=p.y*f;
 } 

func (p * point) Abs() float64{

   return math.Sqrt(p.x*p.x+p.y+p.y)
}

func main(){

	 p:=&point{4,5}
     p.Scale(3)
     fmt.Println(p,p.Abs())

}