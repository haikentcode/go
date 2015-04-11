package main

import (
    
	"fmt"
	"strings"
	"os"
)


func WordCount(s string) map[string]int {
   m :=make(map[string]int);
   a:=strings.Split(s," ")
   
   for i:=0;i<len(a);i++{
   m[a[i]]=m[a[i]]+1
   }
   
  return m

}

func main() {
  
    a:=os.Args[1:]

    var s string;

    for i:=0;i<len(a);i++{

    	s=s+a[i]
    	if(i!=len(a)-1){
    		s=s+" "
    	}
    }

    b:=WordCount(s)
        
    fmt.Println(b)    
    
    for i:=0;i<len(b);i++{

    	fmt.Printf("%s=%d\n",a[i],b[a[i]])
    }
	
}
