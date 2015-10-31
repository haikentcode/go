package main

  import (
           "fmt"
           "log"
           "gopkg.in/mgo.v2"
           "gopkg.in/mgo.v2/bson"
    )

     type Person struct{
              Name string
              Phone string
     }


   func main(){
           session,err :=mgo.Dial("server1.example.com,server2.example.com")
            if err!=nil{
                  panic(err)
             }
             defer session.Close()
            //Optional. Switch the session to a monotonic behavior
            session.SetMode(mgo.Monotonic,true)

            c:=session.DB("test").C("people")
            err=c.Insert(&Person{"aditi","+918442010919"},&Person{"hitesh","+918591216563"})
               
            if err !=nil{
                     log.Fatal(err)
            }

            result :=Person{}
            err =c.Find(bson.M{"name":"aditi"}).One(&result)
            if err !=nil {
                log.Fatal(err)
            }
            fmt.Println("Phone:",result.Phone)
       }    
