package main

import "fmt"
import "github.com/colinmarc/hdfs"

func main() {
    client, err := hdfs.New("localhost:39000")

    if err!=nil{
        fmt.Println(err)
        return
    }
    file, err := client.Open("/1212121.sql")
    fmt.Println("success")
    if err!=nil{
        fmt.Println(err)
        return
    }

    buf := make([]byte, 59)
    file.ReadAt(buf, 48847)

    fmt.Println(string(buf))
    fmt.Println(len(string(buf)))
    fmt.Println("success2")

}

// => Abominable are the tumblers into which he pours his poison.
