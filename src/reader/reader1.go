package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	
	r := strings.NewReader("love me like you do !")

	b := make([]byte, 1)

	for {
	   
		n,err:=r.Read(b)
		fmt.Printf("%s\n",b[:n])
		
		if err == io.EOF {
			break
		}

	}
}
