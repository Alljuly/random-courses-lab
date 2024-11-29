package main

import "fmt"


type user struct{
     name string
     age int8
}

func main(){
    var u user
    fmt.Println(u)



    slice := make([]float32, 10, 12)

    fmt.Println(slice)
    slice = append(slice, 10)
    fmt.Println(slice)
    fmt.Println(len(slice))
    fmt.Println(cap(slice))
    
    fmt.Println("---------")
    fmt.Println(slice)
    slice = append(slice, 10)
    fmt.Println(slice)
    fmt.Println(len(slice))
    fmt.Println(cap(slice))

    
    fmt.Println("---------")
    fmt.Println(slice)
    slice = append(slice, 10)
    fmt.Println(slice)
    fmt.Println(len(slice))
    fmt.Println(cap(slice))
}