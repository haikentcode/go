package main

import "fmt"
import "time"

 
 func counter(s string){

 	for i:=0;i<10;i++{
 		fmt.Println(s,i);
        time.Sleep(time.Second)
 	}
 }


 func main(){


 	go counter("aditi")
 	go counter("hitesh")
    var s string
 	fmt.Scanln(&s)
 	fmt.Println("->",s)

 }