package main
import "fmt"
import "time"

func main() {
	
    msg:=make(chan string)

    go func(){ 

            time.Sleep(time.Second*3)
    	    msg<-"i love you aditi" 

    	}()


     fmt.Println(<-msg)


}