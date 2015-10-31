package main

type linkList struct{
  data int
  next *linkList
}

func Node(data int)(*linkList){

  var tm=make(*linkList)
  (*tm).next=nil
  (*tm).data=data
  
return tm
}


func main(){

var head *linkList

head=Node(4)
(*head).next=Node(6)

  
}
