package main

import (
	"github.com/sandertv/go-raknet"
    "fmt"
)

func main() {
    listener, _ := raknet.Listen("0.0.0.0:19132")
    defer listener.Close()
    for {
        conn, _ := listener.Accept()
        defer conn.Close()
        fmt.Println("Received conn")
        //b := make([]byte, 1024*1024*4)
        b := make([]byte, 3)
        str1, err := conn.Read(b)
        if err != nil{
             fmt.Println("Read error")
         }else{
             fmt.Println(len(b))
             fmt.Println(b)
             fmt.Println(str1)
         }
         str2,err2 := conn.Write([]byte{1, 2, 3, 4})
        if err2 != nil{
             fmt.Println("Write error")
         }else{
             fmt.Println(str2)
         }

    }
}
