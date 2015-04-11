package main

import "fmt"
import "math"

type geometry interface{
  
      area() float64
      perim() float64

}


type square struct{
	height,width float64
}

type circle struct{
	radious float64
}

func (s square) area()float64{

	return s.width*s.height
	
}

func (c circle) area()float64{
	return math.Pi*c.radious*c.radious
}


func (s square) perim()float64 {

	return 2*(s.width+s.height)
	
}

func (c circle) perim()float64{

	return 2*math.Pi*c.radious
}


func calcu(g geometry){

    fmt.Println(g);
    fmt.Println(g.area())
    fmt.Println(g.perim())

}


func main(){
	
	a:=square{3,4}
	b:=circle{4}

	calcu(a)
	calcu(b)
	
}
